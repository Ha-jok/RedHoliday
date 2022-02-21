var jsImg_one = document.getElementById("img_one");
//当前显示的图片
var currentPage_one = 0;
//背景div
var jsBox_one = document.getElementById("box_one");
//左右按钮
var jsLeftBtn_one = document.getElementById("leftbt_one");
var jsRightBtn_one = document.getElementById("rightbt_one");
//创建定时器播放图片
var pTimer_one = setInterval(func_one, 4000);
jsLeftBtn_one.style.display="none";
jsRightBtn_one.style.display="none";
function func_one() {
  currentPage_one++;
  if (currentPage_one == 3) {
    currentPage_one = 0;
  }
  //修改图片
  jsImg_one.src = "images/img" + currentPage_one + ".png";
}

//给box添加鼠标进入事件
jsBox_one.addEventListener("mousemove", function () {
  jsLeftBtn_one.style.display="block";
  jsRightBtn_one.style.display="block";
  //停止定时器 即销毁 所以下方还得在一次启动
  clearInterval(pTimer_one);
}, false);
jsBox_one.addEventListener("mouseout", function () {
  jsLeftBtn_one.style.display="none";
  jsRightBtn_one.style.display="none";
  //启动定时器
  pTimer_one = setInterval(func_one,4000);
}, false);

//左右按钮
jsLeftBtn_one.addEventListener("mouseover", function () {
  jsLeftBtn_one.style.backgroundColor = "rgba(0,0,0,0.6)";
}, false);
jsRightBtn_one.addEventListener("mouseover", function () {
  jsRightBtn_one.style.backgroundColor = "rgba(0,0,0,0.6)";
}, false);

jsLeftBtn_one.addEventListener("mouseout", function () {
  jsLeftBtn_one.style.backgroundColor = "rgba(0,0,0,0.2)";
}, false);
jsRightBtn_one.addEventListener("mouseout", function () {
  jsRightBtn_one.style.backgroundColor = "rgba(0,0,0,0.2)";
}, false);

//鼠标移入显示的左边按钮<
jsLeftBtn_one.addEventListener("click", function () {
  currentPage_one--;
  if (currentPage_one == -1) {
    currentPage_one = 2;
  }
  jsImg_one.src = "images/img" + currentPage_one + ".png";
}, false);

//右边按钮>
jsRightBtn_one.addEventListener("click", function () {
  currentPage_one++;
  if (currentPage_one == 2) {
    currentPage_one = 0;
  }
  jsImg_one.src = "images/img" + currentPage_one + ".png";
}, false);
