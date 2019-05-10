$(document).ready(function () {
    $.ajax({
        url: "/account",
        success: function (result) {
            console.log(result);
            var json = eval("(" + result + ")");
            $("#publicKey").append(json.publicKey);
            $("#privateKey").append(json.privateKey);
            $("#address").append(json.address);
        }
    });
    $(".box_right_sigle").hide();
    $("#count").show();
    $(".hover").click(function(){
        var index=$(".hover").index(this);
        $(".box_right_sigle").hide();
        $(".box_right_sigle").eq(index).show();
    });
    $("#left_myorder").click(function(){
        $.ajax({
            url: "/getmyoder",
            success: function (result) {
                console.log(result);
                console.log(result["id"]);
                var html=' <table>\n' +
                    '               <tr>\n' +
                    '                   <td class="tr_left">商品编号：</td><td>'+result["id"]+'</td>\n' +
                    '               </tr>\n' +
                    '               <tr>\n' +
                    '                   <td class="tr_left">商品名称：</td><td>'+result["name"]+'</td>\n' +
                    '               </tr>\n' +
                    '               <tr>\n' +
                    '                   <td class="tr_left">商品价格：</td><td>'+result["price"]+'</td>\n' +
                    '               </tr>\n' +
                    '               <tr>\n' +
                    '                   <td class="tr_left">商品数量：</td><td>'+result["count"]+'</td>\n' +
                    '               </tr>\n' +
                    '               <tr>\n' +
                    '                   <td class="tr_left">支付总额：</td><td>'+result["allPrice"]+'</td>\n' +
                    '               </tr>\n' +
                    '               <tr>\n' +
                    '                   <td class="tr_left">支付地址：</td><td>'+result["payToAddress"]+'</td>\n' +
                    '               </tr>\n' +
                    '           </table>';
                $("#myorder_order").html(html);

                $("#sure").html('<p id="i_sure">'+result["isSend"]+'</p>');
                $("#i_sure").click(function () {
                    $.ajax({
                        url: "http://localhost:9090/Broadcast",
                        success:function () {
                            alert(result["isSend"]);
                        },
                        error:function () {
                            alert(result["isSend"]);
                        }
                    });
                });
            }
        });
    });
    $("#left_trade").click(function () {
        $.ajax({
            url: "http://localhost:9090/GetBlockChain",
            success: function (result) {
                var html="";
                $.each(result,function (name,value) {
                    var result = eval("(" + value + ")");
                    var data=eval("("+result["Data"]+")");
                    var dataHtml=' <table id="dataTable">\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">支出地址：</td><td>'+data["fromAddress"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">支出金额：</td><td>'+data["money"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">商品信息：</td><td>'+data["shopInfo"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">收入地址：</td><td>'+data["toAddress"]+'</td>\n' +
                        '               </tr>\n' +
                        '           </table>';

                   html+=' <table>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">区块序号：</td><td>'+result["Index"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">时间戳：</td><td>'+result["TimeStamp"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">难度值：</td><td>'+result["Diff"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">前一区块哈希值：</td><td>'+result["PreHash"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">当前区块哈希值：</td><td>'+result["HashCode"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">随机数：</td><td>'+result["Nonce"]+'</td>\n' +
                        '               </tr>\n' +
                        '               <tr>\n' +
                        '                   <td class="tr_left">区块数据：</td><td>'+dataHtml+'</td>\n' +
                        '               </tr>\n' +
                        '           </table>';
                   $("#block").html(html);

                });
            }
        });
    });
});