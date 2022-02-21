// window.onload = function () {
//   var pwd = document.getElementById("pwd");
//   var user = document.getElementById("user");
//   var submit = document.getElementById("submit");
//
//   user.onblur = function () {
//     //用户名正则，4到16位（字母，数字，下划线，减号）
//     var reg_one = /^[a-zA-Z0-9_-]{4,16}$/;
//     if (reg_one.test(this.value)) {
//       alert("格式正确")
//     } else {
//       alert("账号格式不对，4到16位（字母，数字，下划线，减号），到请重新输入");
//     }
//   };
//   pwd.onblur = function () {
//     //密码强度正则，最少6位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符
//     var reg_two = /^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*? ]).*$/;
//     if (reg_two.test(this.value)) {
//       alert("格式正确")
//     } else {
//       alert("密码格式不对，最少6位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符，请重新输入");
//     }
//   };
// };