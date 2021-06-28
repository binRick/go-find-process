package findprocess

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/shirou/gopsutil/process"
)

type ProcessEnvironments map[int64]map[string]string

func (pe *ProcessEnvironments) Pids() *[]int64 {
	pids := []int64{}

	return &pids

}

var PRESHARED_ENV_KEYS = []string{
	`WGCS_SETUP`,
}

var PRESHARED_ENV_BASE64_ENCODED_VALUE_KEYS = []string{
	`WGCS_SETUP`,
}

func env_key_is_encoded(k string) bool {
	for _, kk := range PRESHARED_ENV_BASE64_ENCODED_VALUE_KEYS {
		if strings.ToUpper(kk) == strings.ToUpper(k) {
			return true
		}
	}
	return false
}

func json_dec(s string) *map[string]interface{} {
	dc_dec_json_dec := map[string]interface{}{}
	jerr := json.Unmarshal([]byte(s), &dc_dec_json_dec)

	if jerr == nil {
		return &dc_dec_json_dec

	}
	return &map[string]interface{}{}
}

func b64_dec(s string) string {
	_s := s

	DC_dec, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		_dec, _dec_err := base64.StdEncoding.DecodeString(fmt.Sprintf(`%s=`, s))
		if _dec_err == nil {
			_s = fmt.Sprintf(`%s`, _dec)
		}
	} else {
		_s = fmt.Sprintf(`%s`, DC_dec)
	}
	return _s
}

func Pids() ([]int64, []interface{}, time.Duration, error) {
	pids := []int64{}
	_pids := []int64{}
	pid_matches := map[int64]map[string]string{}
	started := time.Now()
	procs, err := process.Processes()
	Fatal(err)
	unique_val_hashes := map[string]string{}
	unique_vals := []interface{}{}

	for _, P := range procs {
		pids = append(pids, int64(P.Pid))
	}

	for _, pid := range pids {
		pid_env, err := ReadProcessEnvironment(pid)
		if err != nil {
			continue
		}
		pid_env_str := NullTermToStrings(pid_env)
		em := StringsToEnvironmentMap(pid_env_str)
		for k, v := range em {
			for _, pek := range PRESHARED_ENV_KEYS {
				if strings.ToUpper(k) == strings.ToUpper(pek) {
					if env_key_is_encoded(strings.ToUpper(pek)) {
						dec := b64_dec(fmt.Sprintf(`%s=`, v))
						j1 := json_dec(dec)
						j1_str, j1_str_err := json.Marshal(j1)
						if j1_str_err == nil {
							v = fmt.Sprintf(`%s`, j1_str)
						}
					}
					if len(string(v)) > 0 && (len(pek) > 0) {
						_, has := pid_matches[pid]
						if !has {
							pid_matches[pid] = map[string]string{}
							_pids = append(_pids, pid)
						}
						pid_matches[pid][strings.ToUpper(pek)] = string(v)

						hash_str := fmt.Sprintf("%x", md5.Sum([]byte(string(v))))
						_, uv_has := unique_val_hashes[hash_str]
						if !uv_has {
							unique_val_hashes[hash_str] = string(v)
							unique_vals = append(unique_vals, string(v))
						}
					}
				}
			}
		}
	}

	return _pids, unique_vals, time.Since(started), err
}
func Fatal(err error) {
	if err == nil {
		return
	}
	Panic(err)
}

func Panic(e error) {
	if e != nil {
		panic(fmt.Sprintf("%s", e))
	}
}

func NullTermToStrings(b []byte) (s []string) {
	nt := 0
	ntb := byte(nt)
	for {
		i := bytes.IndexByte(b, ntb)
		if i == -1 {
			break
		}
		s = append(s, string(b[0:i]))
		b = b[i+1:]
	}
	return
}

func ReadProcessEnvironment(pid int64) ([]byte, error) {
	proc_path := fmt.Sprintf(`/proc/%d/environ`, pid)
	b, err := ioutil.ReadFile(proc_path)
	if err != nil {
		return b, err
	}
	return b, nil
}

func StringsToEnvironmentMap(s []string) map[string]string {
	em := map[string]string{}
	for _, _s := range s {
		em_s := strings.Split(_s, "=")
		KEY := ``
		VAL := ``
		for ems_s_index, ems_s_v := range em_s {
			if ems_s_index == 0 {
				KEY = ems_s_v
			} else {
				ems_s_v = strings.Replace(ems_s_v, `"`, ``, -1)
				if len(VAL) == 0 {
					VAL = fmt.Sprintf(`%s`, ems_s_v)
				} else {
					VAL = fmt.Sprintf(`%s%s`, VAL, ems_s_v)
				}
			}
		}
		if len(KEY) > 0 && len(VAL) > 0 {
			em[KEY] = VAL
		}
	}
	return em
}
