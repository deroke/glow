package gitprovider

import (
	"regexp"

	"github.com/meinto/glow"
	l "github.com/meinto/glow/logging"
)

type loggingService struct {
	next Service
}

func NewLoggingService(s Service) Service {
	l.Log().Info(l.Fields{"service": s})
	return &loggingService{s}
}

func (s *loggingService) Close(b glow.Branch) (err error) {
	l.Log().Info(l.Fields{"branch": b.BranchName()})
	defer func() {
		l.Log().Error(err)
	}()
	return s.next.Close(b)
}

func (s *loggingService) Publish(b glow.Branch) (err error) {
	l.Log().Info(l.Fields{"branch": b.BranchName()})
	defer func() {
		l.Log().Error(err)
	}()
	return s.next.Publish(b)
}

func (s *loggingService) GetCIBranch() (branch glow.Branch, err error) {
	l.Log().Info(l.Fields{"branch": branch.BranchName()})
	defer func() {
		l.Log().Error(err)
	}()
	return s.next.GetCIBranch()
}

func (s *loggingService) DetectCICDOrigin() (cicdOrigin string, err error) {
	re := regexp.MustCompile(`:([^:]+)@`)

	cleanedCicdOrigin := cicdOrigin
	if re.MatchString(cicdOrigin) {
		cleanedCicdOrigin = re.ReplaceAllString(cicdOrigin, ":xxxxxx@")
	}

	defer func() {
		l.Log().
			Info(l.Fields{"cicdOrigin": cleanedCicdOrigin}).
			Error(err)
	}()
	return s.next.DetectCICDOrigin()
}
