package actions

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/home").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "home page")
}
