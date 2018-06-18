package sav

type BaseApplication struct {

}

func (ctx BaseApplication) Fetch(action Action, handler DataHandler) (Response, error){
  return nil, nil
}

func NewApplication() Application {
  res := &BaseApplication{
  }
  return res
}
