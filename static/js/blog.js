$(document).ready(function() {
	/* ======= Highlight.js Plugin ======= */ 
    /* Ref: https://highlightjs.org/usage/ */     
    $('pre code').each(function(i, block) {
	    hljs.highlightBlock(block);
	 });
	// 自动选中菜单
	let pathname = document.location.pathname;
	$(".navbar-nav li").each(function (index, item) {
		$(item).removeClass('active').find("a>span").remove();
		if (pathname === $.trim($(item).find("a").attr("href"))) {
			console.log(pathname)
			$(item).addClass('active').find("a").append("<span class=\"sr-only\">(current)</span>");
		}
	});
	$("#handle-toc").click(function () {
		$(".editormd-markdown-toc").toggle();
	});
	/**
	 * show msg
	 * @param type
	 * @param message
	 * @param cb callback func
	 */
	function msg(message, type, cb){
		let default_class = "bg-primary";
		if (type === "warning") {
			default_class = "bg-warning";
		} else if (type === "error") {
			default_class = "bg-danger"
		}
		if (message === undefined || message === "") {
			message = "您好! 这是一个Toasts弹出提示框.";
		}
		let tooast = $("#toast");
		tooast.find(".toast-header").removeClass("bg-primary", "bg-warning", "bg-danger").addClass(default_class);
		tooast.find(".toast-body").text(message);
		$(".toast").toast({
			"delay": 30000,
		}).toast("show");
		$('#toast').on('hidden.bs.toast', function () {
			if (typeof cb == "function") {
				cb();
			}
		})
	}
    // login form
	$("#form-login").submit(function(e){
		e.preventDefault();
		let data = $(this).serializeArray();
		$.post("/login", data, function (respoonse) {
			if (respoonse.status > 0) {
				msg(respoonse.msg, "error");
				return;
			}
			window.location.href = "/";
		}, 'json');

		return false;
	});

	// article form
	$("#form-create-article").submit(function (e) {
		e.preventDefault();
		let data = $(this).serializeArray();
		let directory = $("#custom-toc-container").html();
		data.push({
			"name":"directory_html",
			"value":directory
		});
		// 参数判断
		$.post("/article/create", data, function (response) {
			if (response.status > 0) {
				msg(response.msg, "error");
				return;
			}
			// 跳转
			window.location.href = "/detail/"+response.data.ID;
		}, "json")
	});

	// 文档目录
	let directory_content = $("#directory_content");
	directory_content.find("a").on("click", function () {
		event.stopPropagation();
	});
	$("#directory").on("click",function (event) {
		event.stopPropagation();
		directory_content.animate({width:"300px"}, 300, "swing", function () {});
	});
	$("body").click(function (event) {
		$("#directory_content").animate({width:"0", display:"none"}, 300, "swing", function(){});
	});
});