package controllers

import (
	"bufio"
	"ccxt.com/nginxconf/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/session"
	"github.com/codeskyblue/go-sh"
	"io"
	"os"
	"strconv"
	"text/template"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	// _, action := c.GetControllerAndAction()
	// fmt.Println(action)
	// if action == "Index" || action == "Login" {

	// } else {
	// 	if c.GetSession("user") != nil && c.GetSession("pwd") != nil {
	// 		userhttp := c.GetSession("user").(string)
	// 		pwdhttp := c.GetSession("pwd").(string)
	// 		user := beego.AppConfig.String("user")
	// 		pwd := beego.AppConfig.String("pwd")
	// 		if userhttp == user && pwdhttp == pwd {
	// 			//c.Redirect("/domain", 302)
	// 			c.SetSession("user", userhttp)
	// 			c.SetSession("pwd", pwdhttp)
	// 		} else {
	// 			c.Redirect("/", 302)
	// 		}
	// 	} else {
	// 		c.Redirect("/", 302)
	// 	}
	// }

}

func (c *MainController) Index() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplNames = "index.tpl"
}

func (c *MainController) Login() {
	userhttp := c.GetString("user")
	pwdhttp := c.GetString("pwd")
	user := beego.AppConfig.String("user")
	pwd := beego.AppConfig.String("pwd")
	fmt.Println(userhttp)
	fmt.Println(pwdhttp)
	fmt.Println(user)
	fmt.Println(pwd)
	if userhttp == user && pwdhttp == pwd {
		c.Redirect("/domain", 302)
		c.SetSession("user", userhttp)
		c.SetSession("pwd", pwdhttp)
	} else {
		c.Redirect("/", 302)
	}

}

/**虚拟主机列表**/
func (c *MainController) Domain() {
	var (
		domian  models.Domain
		domians []models.Domain
		option  models.Option
		//options []models.Option
	)
	optionerr := option.Query().Filter("name", "confdir").One(&option)
	domian.Query().All(&domians)

	if optionerr == nil {
		c.Data["option"] = option
	} else {

	}

	if len(domians) > 0 {
		c.Data["domians"] = domians
	} else {

	}
	c.TplNames = "domain.tpl"
}

/**设置nginx路径**/
func (c *MainController) Addconfdir() {
	var (
		option models.Option
	)
	//optionerr := option.Query().Filter("name", "confdir").One(&option)
	confdir := c.GetString("confdir")
	_, err := option.Query().Filter("name", "confdir").Update(orm.Params{"value": confdir})
	if err != nil {
		c.Ctx.WriteString("设置失败，错误：" + err.Error())
	}
	c.Ctx.WriteString("设置成功，请返回刷新列表")
}

/**添加虚拟主机 **/
func (c *MainController) Addomain() {
	var (
		domian models.Domain
		//domians []models.Domain
	)
	// domian.Query().All(&domians)
	// if len(domians) > 0 {
	// 	c.Data["domians"] = domians
	// } else {

	// }

	Server_name := c.GetString("Server_name")
	Port, _ := c.GetInt64("Port")
	Proxy_pass := c.GetString("Proxy_pass")
	Access_log := c.GetString("Access_log")
	Expires := c.GetString("Expires")
	Root := c.GetString("Root")

	domian.Access_log = Access_log
	domian.Expires = Expires
	domian.Port = Port
	domian.Proxy_pass = Proxy_pass
	domian.Root = Root
	domian.Server_name = Server_name
	_, err := models.AddDomain(&domian)
	if err != nil {
		c.Ctx.WriteString("添加失败，错误：" + err.Error())
	}
	c.Ctx.WriteString("添加成功，请返回刷新列表")
}

/**虚拟主机列表**/
func (c *MainController) Create() {
	var (
		domian models.Domain
		//tmpdomian models.Domain
		//domians []models.Domain
		option  models.Option
		id      int64
		confdir string
	)

	if id, _ = strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64); id < 1 {
		c.Ctx.WriteString("错误的id")
		//this.RenderString()
		return
	}

	optionerr := option.Query().Filter("name", "confdir").One(&option)
	if optionerr == nil {
		confdir = option.Value
	} else {
		confdir = ""
	}

	domian.Query().Filter("id", id).One(&domian)
	// domian.Access_log
	// domian.Expires
	// domian.Port
	// domian.Proxy_pass
	// domian.Root
	// domian.Server_name
	paththis, _ := os.Getwd()
	d := string(os.PathSeparator)
	fmt.Println(paththis)
	//path = paththis + d + "static" + d + "upload" + d + path + d + h.Filename
	pathdir := paththis + d + "views" + d + "domainconf.tpl"

	s1, tmperr := template.ParseFiles(pathdir)
	if tmperr != nil {
		fmt.Println(tmperr.Error())
	}
	fmt.Println(confdir)
	/**********test write********* /
	file, err := os.Create(confdir + d + domian.Server_name + ".conf")
	if err != nil {
		fmt.Println("writer", err)
		c.Ctx.WriteString("not Create ok")
		return
	}
	defer file.Close()
	//writer := bufio.NewWriter(file)
	/**********test write end********* /
	//s1.Execute(c.Ctx.ResponseWriter, domian)//这个也可以使用
	//domiant := models.Domain{Id: 1, Server_name: "", Port: 80, Proxy_pass: "", Access_log: "", Expires: "", Root: ""}
	exeuteerr := s1.Execute(file, domian)
	if exeuteerr != nil {
		fmt.Println(exeuteerr.Error())
	}

	/****test writer to string****/
	r, w := io.Pipe()
	var tmp string
	tmp = ""
	//data := make([]byte, 2048)
	//r.Read(data)

	data := make([]byte, 1024)

	go func() {
		for n, err := r.Read(data); err == nil; n, err = r.Read(data) {
			fmt.Printf("%s", data[:n])
			tmp = tmp + string(data[:n])
		}
		// for {
		// 	r.Read(data)
		// 	tmp = tmp + string(data)
		// 	//fmt.Println(string(b[0:])) //hello widuu
		// }
	}()

	//var b [128]byte
	// go func() {
	// 	for {
	// 		r.Read(b[0:])
	// 		tmp = tmp + string(b[0:])
	// 		//fmt.Println(string(b[0:])) //hello widuu
	// 	}

	// }()

	s1.Execute(w, domian)
	fmt.Println(tmp)
	//fmt.Println("read number", n)

	/********/

	csh := sh.Command("/etc/init.d/nginx", "reload")
	csh.Start()
	csh.Run()
	msgcsh, _ := csh.Output()
	fmt.Println(msgcsh)

	msg, _ := sh.Command("/etc/init.d/nginx", "reload").Output()

	c.Ctx.WriteString(string(msg))

	c.Ctx.WriteString("ok")
	//writer.Flush()

	//this.RenderString()
	return
	//c.TplNames = "domain.tpl"
}

/**删除虚拟主机**/
func (c *MainController) Delete() {
	var (
		domian models.Domain
		id     int64
	)

	if id, _ = strconv.ParseInt(c.Ctx.Input.Param(":id"), 10, 64); id < 1 {
		c.Ctx.WriteString("错误的id")
		//this.RenderString()
		return
	}

	_, err := domian.Query().Filter("id", id).Delete()
	if err != nil {
		c.Ctx.WriteString("删除失败，错误：" + err.Error())
	}
	c.Ctx.WriteString("删除成功，请返回刷新列表")
}

func writeResult(vals []int, outfile string) error {

	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("writer", err)
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, v := range vals {

		writer.WriteString(strconv.Itoa(v))
		writer.WriteString("\n")
		writer.Flush()
	}

	return err
}

func (c *MainController) Tes() {
	// test := beego.GlobalSessions
	// test.SessionStart(w, r)
	// test()
	c.Ctx.WriteString("ok")
}

func test() {
	// var (
	// 	con1 beego.Controller
	// 	con2 *beego.Controller
	// )
	// beego.
	// 	con1.SetSession("name", "value")
	// con2.SetSession("name", "value")
	// test2 := con1.GetSession("name")
	// //con1 = new(beego.Controller)
	// test := con2.GetSession("name")
	// fmt.Println(test)
	// fmt.Println(test2)

}
