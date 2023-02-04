package localizator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var globMapLocales sync.Map

func setToglobMapLocales(locale, space string, m map[string]interface{}) {
	globMapLocales.Store(locale+"."+space, m)
}

func ParseDir(path string, space string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename := file.Name()
		tag, err := language.Parse(filename[0 : len(filename)-len(filepath.Ext(filename))])
		if err != nil {
			return fmt.Errorf("Parsing files. Filename %s, error: %v", filename, err)
		}

		b, err := ioutil.ReadFile(filepath.Join(path, filename))
		if err != nil {
			return fmt.Errorf("Reading file. Filename %s, error: %v", filename, err)
		}

		locales := make(map[string]interface{})
		if err = json.Unmarshal(b, &locales); err != nil {
			return fmt.Errorf("Serialization file. Filename %s, error: %v", filename, err)
		}

		setToglobMapLocales(tag.String(), space, locales)
		for k, v := range locales {
			if space != "" {
				k = space + "." + k
			}

			if ss, ok := v.(string); ok {
				message.SetString(tag, k, ss)

			}
		}
	}

	return nil
}

func GetKey(locale, space, key string) string {
	if space != "" {
		key = space + "." + key
	}
	tag := message.MatchLanguage(locale, "en")
	return message.NewPrinter(tag).Sprintf(message.Key(key, ""))
}
