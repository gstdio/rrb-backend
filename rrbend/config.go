package rrbend

import (
	"fmt"
	"os"
	"time"
	//"strings"
	"github.com/gwtony/gapi/config"
	"github.com/gwtony/gapi/errors"
)

type RrbendConfig struct {
	adminToken    string        /* admin token */

	apiLoc        string        /* rrbend api location */

	maddr         string        /* mysql addr */
	dbname        string        /* db name */
	dbuser        string        /* db username */
	dbpwd         string        /* db password */

	retention     time.Duration /* retention time in backend */

	timeout       time.Duration
}

// ParseConfig parses config
func (conf *RrbendConfig) ParseConfig(cf *config.Config) error {
	var err error
	if cf.C == nil {
		return errors.BadConfigError
	}
	conf.maddr, err = cf.C.GetString("rrbend", "mysql_addr")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] [rrbend] Read conf: No mysql_addr")
		return err
	}
	conf.dbname, err = cf.C.GetString("rrbend", "mysql_dbname")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] [rrbend] Read conf: No mysql_dbname")
		return err
	}
	conf.dbuser, err = cf.C.GetString("rrbend", "mysql_dbuser")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] [rrbend] Read conf: No mysql_dbuser")
		return err
	}
	conf.dbpwd, err = cf.C.GetString("rrbend", "mysql_dbpwd")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Error] [rrbend] Read conf: No mysql_dbpwd")
		return err
	}

	conf.apiLoc, err = cf.C.GetString("rrbend", "api_location")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Info] [rrbend] Read conf: No api_location, use default location:", RRBEND_LOC)
		conf.apiLoc = RRBEND_LOC
	}

	conf.adminToken, err = cf.C.GetString("rrbend", "admin_token")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Info] [rrbend] Read conf: No admin_token, use default admin token:", DEFAULT_ADMIN_TOKEN)
		conf.adminToken = DEFAULT_ADMIN_TOKEN
	}

	rtstr, err := cf.C.GetString("rrbend", "retention_time")
	if err != nil {
		fmt.Fprintln(os.Stderr, "[Info] [rrbend] Read conf: No retention_time, use default retention time:", DEFAULT_RETENTION_TIME)
		conf.retention, _ = time.ParseDuration(DEFAULT_RETENTION_TIME)
	} else {
		retention, err := time.ParseDuration(rtstr)
		if err != nil {
			fmt.Fprintln(os.Stderr, "[Info] [rrbend] Read conf: Parse retention_time failed:", err)
			return err
		}
		fmt.Fprintln(os.Stderr, "[Info] [rrbend] Read conf: Set retention_time to", retention)
		conf.retention = retention
	}

	return nil
}
