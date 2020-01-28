package cron

import (
	"bufio"
	"fmt"
	"github.com/meklis/http-snmpwalk-proxy/logger"
	"github.com/mileusna/crontab"
	"os"
	"os/exec"
	"strings"
	"time"
)

type CronConfig struct {
	PrintFormat string `yaml:"print_format"`
	Jobs        []Job  `yaml:"jobs"`
}

type Job struct {
	Name        string `yaml:"name"`
	Crontab     string `yaml:"crontab"`
	Exec        string `yaml:"job"`
	PrintOutput bool   `yaml:"print_output"`
}

type Cron struct {
	jobs        []Job
	printFormat string
	lg          *logger.Logger
	cron        *crontab.Crontab
}

var (
	lg *logger.Logger
)

func Init(conf CronConfig, log *logger.Logger) (c *Cron) {
	c = new(Cron)
	c.jobs = conf.Jobs
	c.printFormat = conf.PrintFormat
	c.lg = log
	lg = log
	c.cron = crontab.New()
	return c
}

func (c *Cron) Run() {
	c.lg.InfoF("Crontab initialized")
	for _, job := range c.jobs {
		if strings.Trim(job.Crontab, " ") == "@reboot" {
			c.lg.InfoF("Trying adding job '%v'", job.Name)
			go execJob(job.Name, job.Exec, c.printFormat, job.PrintOutput)
			continue
		}
		c.lg.InfoF("Trying adding job '%v'", job.Name)
		if err := c.cron.AddJob(job.Crontab, execJob, job.Name, job.Exec, c.printFormat, job.PrintOutput); err != nil {
			c.lg.CriticalF("Error add job with name %v", job.Name)
			fmt.Println(err)
			time.Sleep(time.Second * 5)
			os.Exit(1)
		}
	}
	for {
		time.Sleep(time.Second * 5)
	}
}
func execJob(name, command, printFormat string, printOutput bool) {
	lg.NoticeF("[%v] start execute command '%v'", name, command)
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Env = os.Environ()
	out, _ := cmd.StdoutPipe()
	rd_out := bufio.NewReader(out)
	err, _ := cmd.StderrPipe()
	rd_err := bufio.NewReader(err)
	go func() {
		time.Sleep(time.Millisecond * 5)
		for cmd != nil {
			str, err := rd_out.ReadString('\n')
			if err != nil {
				return
			}
			str = strings.Trim(str, "\n")
			if printOutput {
				if printFormat == "log" {
					lg.InfoF("[%v] > %v", name, str)
				} else {
					fmt.Printf("[%v] > %v\n", name, str)
				}
			}
		}
	}()
	go func() {
		time.Sleep(time.Millisecond * 5)
		for cmd != nil {
			str, err := rd_err.ReadString('\n')
			if err != nil {
				return
			}
			str = strings.Trim(str, "\n")
			if printOutput {
				if printFormat == "log" {
					lg.InfoF("[%v] > %v", name, str)
				} else {
					fmt.Printf("[%v] > %v\n", name, str)
				}
			}
		}
	}()
	if err := cmd.Run(); err != nil {
		lg.Errorf("[%v] problem execute: %v", name, err.Error())
	}
	code := cmd.ProcessState.ExitCode()
	if code != 0 {
		lg.WarningF("[%v] finished with code %v", name, code)
	} else {
		lg.NoticeF("[%v] finished with code %v", name, code)
	}
}
