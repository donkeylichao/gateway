$(function(){
	waterfall();
});
$(window).resize(function(){
	waterfall();
});
function waterfall(){
	var $boxs=$(".flow-box"); //获取所有的box
	var w=$boxs.eq(0).outerWidth()+10; //获取一个box的宽度
	var cols=Math.floor($("#listBox").width()/w); //求出当前浏览器宽度下可以放几列
	//console.log(cols);
	$("#listBox").css("margin","0 auto"); //设置列表box居中
	//console.log($("#listBox").width());
	var hArr=[]; //数组放高度
	$boxs.each(function(index,value){
		var h=$boxs.eq(index).outerHeight();
		if(index<cols) {
			hArr[index]=h;
			$(value).css({
				'position':'relative',
				'top':'0px',
				'left':'0px'
			});
		}else{
			var minH=Math.min.apply(null,hArr);
			var minHIndex=$.inArray(minH,hArr);
			$(value).css({
				'position':'absolute',
				'top':minH+10+'px',
				'left':minHIndex*w+'px'
			});
			hArr[minHIndex]+=$boxs.eq(index).outerHeight()+11;
		}
		var maxH=Math.max.apply(null,hArr);
		$("#listBox").css("height",maxH+80+"px");
	})
};
//flow-box阴影效果
$(".flow-box").hover(function(){
	$(this).addClass('shadow');
},
function(){
	$(this).removeClass('shadow');
});

//导航条的active效果
$(".navbar-nav li").click(function(){
	$(".navbar-nav li").removeClass('active');
	$(this).addClass('active');
});
