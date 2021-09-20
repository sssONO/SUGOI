package entity

//DB上のテーブル、カラムと構造体との関連付けが自動的に行われる
type Product struct {
	ID   int
	Name string `json:"name"`
}
