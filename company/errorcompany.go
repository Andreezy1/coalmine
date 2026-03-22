package company

import "errors"

var ErrorNotMoney = errors.New("У вас недостаточно денег!")
var ErrorHaveItem = errors.New("У вас уже есть этот предмет!")
var ErrorCloseGame = errors.New("Вы еще не все купили!")
