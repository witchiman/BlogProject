{{define "navbar"}}
<header class="navbar  navbar-static-top">
    <div class="container">
        <div class="navbar-header">
            <a class="navbar-brand" href="/">我的博客</a>
        </div>
        <nav class="collapse navbar-collapse">
            <ul class="nav navbar-nav">
                <li {{if .IsHome}} class="active" {{end}}><a href="/">首页</a></li>
                <li {{if .IsCategory}} class="active" {{end}}><a href="/categories">分类</a></li>
                <li {{if .IsTopic}} class="active" {{end}}><a href="/topic">文章</a></li>
            </ul>
        </nav>
        <nav class=" pull-right">
            <ul class="nav navbar-nav">
                {{if .IsLogin}}
                <li><a href="/login?exit=true">退出登录</a></li>
                {{else}}
                <li><a href="/login">登录</a></li>
                {{end}}
            </ul>
        </nav>
    </div>

</header>

{{end}}