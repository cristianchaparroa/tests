package main


type Charater struct {
    index int
    repeated bool
    s string
}

func firstNotRepeatingCharacter(s string) string {
  m:= make(map[string]*Charater,0)

  for i := 0; i<len(s); i++ {
    v := s[i]

    key := string(v)

    val, ok := m[key]

    if ok {
      val.repeated = true
    } else {
      m[string(v)] = &Charater{repeated:false, index:i, s:key}
    }
  }

  fnri := 100000
  fnrc := ""

  for _, v := range m {
    i := v.index
    repeated := v.repeated

    if (repeated == false && i < fnri) {
      fnri = i
      fnrc = v.s
    }
  }

  if (fnrc != "") {
    return fnrc
  }

  return "_"
}
