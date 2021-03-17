var now = timestamp = Date.parse(new Date()) / 1000;
var expire = 0;
var sign_obj = '';
var g_object_name = "";
var infos;
var file = document.getElementById('uploadImg');
file.onchange = function() {
    var fileData = this.files[0];//这是我们上传的文件
    fileUpload(fileData);
};
getListImage();
function getListImage() {
    $.get("/api/list_img", {}, function(res) {
        if(res.code == 200){
            //教师列表数据
            var listData = res.data;
            var str = '<div id="fh5co-board" data-columns="4">';
            if (listData) {
                var number = listData.length;
                var one = '', two = '', three = '', four = '';
                $.each(listData,function(index,value){
                    var num  = (index + 1) % 4;
                    if (num === 1) {
                        if (!one) {
                            one += '<div class="column size-1of4">';
                        }
                        one += '<div class="item">\n' +
                            '       <div class="animate-box bounceIn animated">\n' +
                            '           <a href="./img/web/home/img_1.jpg" class="image-popup fh5co-board-img"\n' +
                            '                               title="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Explicabo, eos?">\n' +
                            '                 <img src="' + value.true_url + '" alt="' + value.title + '">\n' +
                            '             </a>\n' +
                            '        </div>\n' +
                            '       <div class="fh5co-desc">' + value.title + '</div>\n' +
                            '    </div>\n';

                    }else if (num === 2) {
                        if (!two) {
                            two += '<div class="column size-1of4">';
                        }
                        two += '<div class="item">\n' +
                            '       <div class="animate-box bounceIn animated">\n' +
                            '           <a href="./img/web/home/img_1.jpg" class="image-popup fh5co-board-img"\n' +
                            '                               title="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Explicabo, eos?">\n' +
                            '                 <img src="' + value.true_url + '" alt="' + value.title + '">\n' +
                            '             </a>\n' +
                            '        </div>\n' +
                            '       <div class="fh5co-desc">' + value.title + '</div>\n' +
                            '    </div>\n';
                    }else if (num === 3) {
                        if (!three) {
                            three += '<div class="column size-1of4">';
                        }
                        three += '<div class="item">\n' +
                            '       <div class="animate-box bounceIn animated">\n' +
                            '           <a href="./img/web/home/img_1.jpg" class="image-popup fh5co-board-img"\n' +
                            '                               title="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Explicabo, eos?">\n' +
                            '                 <img src="' + value.true_url + '" alt="' + value.title + '">\n' +
                            '             </a>\n' +
                            '        </div>\n' +
                            '       <div class="fh5co-desc">' + value.title + '</div>\n' +
                            '    </div>\n';
                    }else if (num === 0) {
                        if (!four) {
                            four += '<div class="column size-1of4">';
                        }
                        four += '<div class="item">\n' +
                            '       <div class="animate-box bounceIn animated">\n' +
                            '           <a href="./img/web/home/img_1.jpg" class="image-popup fh5co-board-img"\n' +
                            '                               title="Lorem ipsum dolor sit amet, consectetur adipisicing elit. Explicabo, eos?">\n' +
                            '                 <img src="' + value.true_url + '" alt="' + value.title + '">\n' +
                            '             </a>\n' +
                            '        </div>\n' +
                            '       <div class="fh5co-desc">' + value.title + '</div>\n' +
                            '    </div>\n';
                    }
                });
                str +=  one + '</div>' + two + '</div>' + three + '</div>' + four + '</div>';
            }
            str += ' </div>';
            $('.row').append(str)
        }
    });
}

function fileUpload(file){
    get_signature();
    if(!sign_obj) {
        layer.msg('系统繁忙，请稍后', {icon: 2, time: 2000});
        return false;
    }
    g_object_name =  sign_obj.dir + guid() + get_suffix(file.name);  //文件名
    var request = new FormData();
    request.append("OSSAccessKeyId",sign_obj.accessid); //Bucket 拥有者的Access Key Id。
    request.append("policy",sign_obj.policy);             //policy规定了请求的表单域的合法性
    request.append("Signature",sign_obj.signature); //根据Access Key Secret和policy计算的签名信息，OSS验证该签名信息从而验证该Post请求的合法性
    request.append("key",g_object_name);            //文件名字，可设置路径
    request.append("success_action_status",'200');  // 让服务端返回200,不然，默认会返回204
    request.append('x-oss-object-acl', 'public-read');
    request.append('file', file);
    $.ajax({
        url : sign_obj.host,  //上传阿里地址
        data : request,
        processData: false,//默认true，设置为 false，不需要进行序列化处理
        cache: false,//设置为false将不会从浏览器缓存中加载请求信息
        async: false,//发送同步请求
        contentType: false,//避免服务器不能正常解析文件---------具体的可以查下这些参数的含义
        dataType: 'xml',//不涉及跨域  写json即可
        type : 'post',
        success : function(callbackHost, request) {     //callbackHost：success,request中就是 回调的一些信息，包括状态码什么的
            infos = {
                oldname:file['name'],
                filesize:file['size'],
                type:file['type'],
                file:g_object_name,
            };
            console.log(infos);
        },
        error : function(returndata) {
            console.log("return data:"+returndata);
            layer.msg('系统出现异常，请稍后', {icon: 2, time: 2000});
        }
    });
    var request = new FormData();
    request.append("img_url",g_object_name);
    request.append("title",file['name']);
    layer.load(2);
    $.ajax({
        url : '/api/upload_img',  //上传阿里地址
        data : request,
        processData: false,//默认true，设置为 false，不需要进行序列化处理
        cache: false,//设置为false将不会从浏览器缓存中加载请求信息
        async: false,//发送同步请求
        contentType: false,//避免服务器不能正常解析文件---------具体的可以查下这些参数的含义
        dataType: 'json',//不涉及跨域  写json即可
        type : 'post',
        success : function(callbackHost, request) {     //callbackHost：success,request中就是 回调的一些信息，包括状态码什么的
            console.log(request);
            layer.close(loading);
        },
    });
    window.location.reload();
    return JSON.stringify(infos);

}

/*获取签名信息验证*/
function get_signature(){
    //可以判断当前expire是否超过了当前时间,如果超过了当前时间,就重新取一下。3s 做为缓冲
    var  now = timestamp = Date.parse(new Date()) / 1000;
    if (expire < now + 3){
        var body = send_request();
        var obj =JSON.parse(body);
        console.log(obj);
        if(obj.code === 200){
            sign_obj= obj.data;
            expire= parseInt(sign_obj['expire']);
            return true;
        }
        return true;
    }
    return false;
}

/**
 * @remark 获取上传的sign
 * @Time 19-8-23 16:03
 */
function send_request(){
    var xmlhttp = null;
    if (window.XMLHttpRequest)
    {
        xmlhttp=new XMLHttpRequest();
    }
    else if (window.ActiveXObject)
    {
        xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
    if (xmlhttp!=null)
    {
        var serverUrl = '/api/oss_policy';
        xmlhttp.open( "GET", serverUrl, false );
        xmlhttp.send( null );
        return xmlhttp.responseText ;
    }else{
        alert("您的浏览器不支持XMLHTTP，请换个浏览器上传");
    }
}

/**
 * @remark 获取文件后缀名
 * @function get_suffix
 * @Time  19-8-23 16:09
 */
function get_suffix(filename) {
    var pos = filename.lastIndexOf('.');
    var suffix = '';
    if (-1 !== pos) {
        suffix = filename.substring(pos);
    }
    return suffix;
}

/**
 * get  guid
 */
function guid() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function(c) {
        var r = Math.random()*16|0, v = c === 'x' ? r : (r&0x3|0x8);
        return v.toString(16);
    });
}
