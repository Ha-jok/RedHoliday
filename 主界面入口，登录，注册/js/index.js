var jsLis = document.getElementsByClassName("li");
//img节点
var jsImg = document.getElementById("img");
//当前显示的图片
var currentPage = 0;
//背景div
var jsBox = document.getElementById("box");
//左右按钮
var jsLeftBtn = document.getElementById("leftbt");
var jsRightBtn = document.getElementById("rightbt");
//第一个小圆点初始化成红色
jsLis[0].style.backgroundColor = "white";
//创建定时器播放图片
var pTimer = setInterval(func, 2000);

function func() {
  currentPage++;
  if (currentPage == 8) {
    currentPage = 0;
  }
  //修改图片
  jsImg.src = "images/" + currentPage + ".jpg";
  //清空所有小圆点的颜色
  for (var i = 0; i < jsLis.length; i++) {
    jsLis[i].style.backgroundColor = "rgba(255,255,255,.4)";
  }
  //修改小圆点的颜色
  jsLis[currentPage].style.backgroundColor = "white";
}

//给box添加鼠标进入事件
jsBox.addEventListener("mousemove", function () {
  //停止定时器 即销毁 所以下方还得在一次启动
  clearInterval(pTimer);
}, false);
jsBox.addEventListener("mouseout", function () {
  //启动定时器
  pTimer = setInterval(func, 2000);
}, false);

//左右按钮
jsLeftBtn.addEventListener("mouseover", function () {
  jsLeftBtn.style.backgroundColor = "rgba(0,0,0,0.6)";
}, false);
jsRightBtn.addEventListener("mouseover", function () {
  jsRightBtn.style.backgroundColor = "rgba(0,0,0,0.6)";
}, false);

jsLeftBtn.addEventListener("mouseout", function () {
  jsLeftBtn.style.backgroundColor = "rgba(0,0,0,0.2)";
}, false);
jsRightBtn.addEventListener("mouseout", function () {
  jsRightBtn.style.backgroundColor = "rgba(0,0,0,0.2)";
}, false);

//鼠标移入显示的左边按钮<
jsLeftBtn.addEventListener("click", function () {
  currentPage--;
  if (currentPage == -1) {
    currentPage = 7;
  }
  jsImg.src = "images/" + currentPage + ".jpg";
  //清空所有小圆点的颜色
  for (var i = 0; i < jsLis.length; i++) {
    jsLis[i].style.backgroundColor = "rgba(255,255,255,.4)";

  }
  //修改小圆点的元素
  jsLis[currentPage].style.backgroundColor = "white";
}, false);

//右边按钮>
jsRightBtn.addEventListener("click", function () {
  currentPage++;
  if (currentPage === 8) {
    currentPage = 0;
  }
  jsImg.src = "images/" + currentPage + ".jpg";
  for (var i = 0; i < jsLis.length; i++) {
    jsLis[i].style.backgroundColor = "rgba(255,255,255,.4)";
  }
  jsLis[currentPage].style.backgroundColor = "white";
}, false);


//每一个小圆点都添加鼠标移入事件
for (var i = 0; i < jsLis.length; i++) {
//此处使用闭包，因为由于闭包会携带包含它的函数的作用域，因此会比其他函数占用更多的内存。过度使用闭包可能会导致内存占用过多。对于性能优化方面，我推荐使用方法一
  jsLis[i].onmouseover = (function (x) {
    function m() {
      currentPage = x;
      jsImg.src = "images/" + currentPage + ".jpg";
      for (var i = 0; i < jsLis.length; i++) {
        jsLis[i].style.backgroundColor = "rgba(255,255,255,.4)";
      }
      jsLis[currentPage].style.backgroundColor = "white";
    }

    return m;
  })(i);
}
