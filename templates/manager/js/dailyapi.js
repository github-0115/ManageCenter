
var defaults = {
    chartUrl: "/api/porn/chart",
    starUserUrl: "/console_detail",
    timechartUrl: "/api/porn/timechart",
    searchUrl:"/stat_search_users",

};
var dom = document.getElementById("daily");
var myChart = echarts.init(dom);
var pieChart = echarts.init(document.getElementById('pie'));
var pPieChart = echarts.init(document.getElementById('p-pie'));
var sPieChart = echarts.init(document.getElementById('s-pie'));
var nPieChart = echarts.init(document.getElementById('n-pie'));
var app = {};
var startTime
var endTime
var cnResult= {
    p:"色情",
    s:"性感",
    n:"正常",
};

var enResult = {
    色情: "p",
    性感: "s",
    正常: "n",
};

var enScreenResult = {
    拍屏: "isScreen",
    非拍屏: "notScreen",
};

var curStatus = "identify"
var curMagnitude = "day"
var username=""
var isAll=true
if (getUrlParam("username") == null) {
    username =""
} else {
    username = getUrlParam("username");
    isAll=false
    $("#lblUser").text("当前查看"+username+"的数据")
}

function getUrlParam(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg);  //匹配目标参数
    if (r != null) return unescape(r[2]);
    return null; //返回参数值
}
option = {
    title: {
        text: '近1天api调用概况',
    },
    tooltip: {
        trigger: 'axis'
    },
    legend: {
        data: ['次数']
    },
    toolbox: {
        show: true,
        feature: {
            dataZoom: {
                show: true
            },
            dataView: {
                readOnly: false
            },
            magicType: {
                type: ['line', 'bar']
            },
            restore: {},
            saveAsImage: {}
        }
    },

    xAxis: {
        type: 'category',
        boundaryGap: false,
        data: []
    },
    yAxis: {
        axisLabel: {
            formatter: '{value} 次'
        }
    },
    series: [{
        name: '鉴黄总量',
        type: 'line',
        data: [],
        barWidth: '20%',

        markLine: {
            data: [{
                type: 'average',
                name: '平均值'
            },
            ]
        }
    }],

};

var dataStyle = {
    normal: {
        label: {show: false},
        labelLine: {show: false}
    }
};
pieOption = {
    title: {
        text: '总量:',
        x: '13.5%',
        y: 'center',
        itemGap: 20,
        textStyle: {
            color: 'rgba(30,144,255,0.8)',
            fontFamily: '微软雅黑',
            fontSize: 20,
            fontWeight: 'bolder'
        }
    },
    tooltip: {
        show: true,
        formatter: formatBigPieTip
    },
    color: ['green', 'orange', '#C24641'],
    legend: {
        orient: 'vertical',
        x: document.getElementById('pie').offsetWidth / 2,
        y: document.getElementById('pie').offsetHeight / 1.7,
        itemGap: 12,
        textStyle: {
            fontSize: 16,
        },
        data: ['正常', '性感', '色情'],
        formatter: function (name) {
            return name;
        }
    },
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: false},
            restore: {show: true},
            saveAsImage: {show: true}
        }
    },
    series: [
        {
            name: '总量',
            type: 'pie',
            clockWise: false,
            radius: [90, 140],
            itemStyle: dataStyle,
            center: ['20%', '50%'],
            data: []
        }
    ]
};
pPieOption = {
    title: {
        text: '色情',
        x: 'center',
        y: 'center',
        itemGap: 20,
        textStyle: {
            color: 'rgba(30,144,255,0.8)',
            fontFamily: '微软雅黑',
            fontSize: 15,
            fontWeight: 'bolder'
        }
    },
    tooltip: {
        show: true,
        formatter: formatPieTip
    },
    color: ['#C24641', 'gray'],
    series: [
        {
            name: '色情',
            type: 'pie',
            clockWise: false,
            radius: [25, 45],
            itemStyle: {
                normal: {
                    label: {
                        formatter: formatLabel
                    },
                    labelLine: {length: 0}
                },
            },
        }
    ]
};

sPieOption = {
    title: {
        text: '性感',
        x: 'center',
        y: 'center',
        itemGap: 20,
        textStyle: {
            color: 'rgba(30,144,255,0.8)',
            fontFamily: '微软雅黑',
            fontSize: 15,
            fontWeight: 'bolder'
        }
    },
    tooltip: {
        show: true,
        formatter: formatPieTip
    },
    color: ['orange', 'gray'],
    series: [
        {
            name: '性感',
            type: 'pie',
            clockWise: false,
            radius: [25, 45],
            itemStyle: {
                normal: {
                    label: {
                        formatter: formatLabel
                    },
                    labelLine: {length: 0}
                },
            },
        }
    ]
};

nPieOption = {
    title: {
        text: '正常',
        x: 'center',
        y: 'center',
        itemGap: 20,
        textStyle: {
            color: 'rgba(30,144,255,0.8)',
            fontFamily: '微软雅黑',
            fontSize: 15,
            fontWeight: 'bolder'
        }
    },
    tooltip: {
        show: true,
        formatter: formatPieTip
    },
    color: ['green', 'gray'],
    series: [
        {
            name: '正常',
            type: 'pie',
            clockWise: false,
            radius: [25, 45],
            itemStyle: {
                normal: {
                    label: {
                        formatter: formatLabel
                    },
                    labelLine: {length: 0}
                },
            },
        }
    ]
};

$("#search").click(function(data){
    var searchtext = $("#autocomplete-search").val()
    $.ajax({
        type: 'get',
        url: defaults.searchUrl,
        data:{
            query: searchtext,
            page:"1",
            rows:"10",

        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
            window.location.href="/searchuser?url=datastat&text="+$("#autocomplete-search").val()
        },
        error: function (data) {
            //window.location.href="/searchuser "
        },
    });
});

$('#autocomplete-search').autocomplete({
       serviceUrl: '/stat_search_users',
        params:{
            page:"1",
            rows:"10",
        },
        ajaxSettings:{
            headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
            },
        },
        onSelect: function (suggestion) {
           username=suggestion.data
           isAll=false
           $("#lblUser").text("当前查看"+username+"的数据")
            if(curStatus=="identify"){
                if(curMagnitude=="custom"){
                    getData()
                }else{
                clickIdentify()
                }
            } else{
                if(curMagnitude=="custom"){
                    getData()
                }else{
                    clickScreen()
                }
    }
        },
        transformResult: function(response) {
             response=$.parseJSON(response);
             var result=[]
             response.result.forEach(function(item,idx){
                var resultTemp={}
                resultTemp={
                    "username":item.username,
                    "company":item.company
                }
                result.push(resultTemp)
                if(item.company==""){
                    resultTemp={
                        "username":item.username,
                        "company":item.username
                    }
                    result.push(resultTemp)
                }
             })
                return {
            suggestions: $.map(result, function(dataItem) {
                return { value: dataItem.company, data: dataItem.username};
            })
        }
        }
    });
Date.prototype.format = function (fmt) {
        var o = {
            "M+": this.getMonth() + 1,                 //月份
            "d+": this.getDate(),                    //日
            "h+": this.getHours(),                   //小时
            "m+": this.getMinutes(),                 //分
            "s+": this.getSeconds(),                 //秒
            "q+": Math.floor((this.getMonth() + 3) / 3), //季度
            "S": this.getMilliseconds()             //毫秒
        };
        if (/(y+)/.test(fmt))
            fmt = fmt.replace(RegExp.$1, (this.getFullYear() + "").substr(4 - RegExp.$1.length));
        for (var k in o)
            if (new RegExp("(" + k + ")").test(fmt))
                fmt = fmt.replace(RegExp.$1, (RegExp.$1.length == 1) ? (o[k]) : (("00" + o[k]).substr(("" + o[k]).length)));
        return fmt;
}
function addDate(curDate,curMagnitude){
    curDate = curDate.valueOf()
    curDate = curDate + curMagnitude * 24 * 60 * 60 * 1000
    curDate = new Date(curDate)
return curDate;
}
function formatLabel(params){
    var res = "";
    res = params.name + " : " + params.value.toLocaleString();
    return res;
};

function formatPieTip(params) {
    var res = "";
    res = params.seriesName + "<br/>" + params.name + " : " + params.value.toLocaleString() + " (" + params.percent + "%)"
    return res;
}
function formatBigPieTip(params) {
    var res = "";
    res = params.name + "总量: " + params.value.toLocaleString() + " (" + params.percent + "%)"
    return res;
}

var resSummary = function (data, params, res, type, idx) {
    res += cnResult[type] + ':' + data[type][params[idx].dataIndex].toLocaleString() + '&nbsp;&nbsp;&nbsp;&nbsp;' + '(需复审：' + data[type + '_review'][params[idx].dataIndex].toLocaleString() + '）<br/>';
    return res
};

var sumTotal = function (data) {
    var total = 0;
    data.forEach(function (item, idx) {
        total = total + item
    });
    return total
};

var formatResult = function (data, name) {
    var rate = 0;
    if (sumTotal(data[enResult[name]]) != 0)
        rate = (sumTotal(data[enResult[name] + '_review']) / sumTotal(data[enResult[name]]) * 100).toFixed(2);
    return name + " " + sumTotal(data[enResult[name]]).toLocaleString() + ",复审 " + sumTotal(data[enResult[name] + '_review']).toLocaleString() + "(占比 " + rate + "%)";
}

var formatScreenResult = function (data, name) {
    var rate = 0;
    var total = sumTotal(data.isScreen) + sumTotal(data.notScreen)
    if (total != 0)
        rate = (sumTotal(data[enScreenResult[name]]) / total * 100).toFixed(2);
    return name + " " + sumTotal(data[enScreenResult[name]]).toLocaleString() + "(占比 " + rate + "%)";
}

function screenTotal(isScreen, notScreen) {
    var total = []
    var totalTemp = 0
    isScreen.forEach(function (item, idx) {
        totalTemp = item + notScreen[idx]
        total.push(totalTemp)
    });
    return total
}
 $.ajax({
        type: 'get',
        url: defaults.starUserUrl,
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
        },
        success: function (data) {
            data.startUserList.forEach(function(item,idx){
                var html= '<option value=\"'+item+'\">'+item+'</option>'
                $("#selectPointOfUser").append(html)
            })
            addEven()
        },
        error: function (data) {
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },

    });
initData()
function addEven(){
    $("#selectPointOfUser").change(function(){
        var value=$("#selectPointOfUser").prop("value");
        username=value
        if(value=="all")
            value="全部"
        $("#lblUser").text("当前查看"+value+"的数据")
        if(curStatus=="identify"){
            if(value=="全部"){
                isAll=true
            }else{
                isAll=false
            }
            if(curMagnitude=="custom"){
                getData()
            }else{
                clickIdentify()
            }
    } else{
        if(value=="all"){
                isAll=true
            }else{
                isAll=false
            }
            if(curMagnitude=="custom"){
                getData()
            }else{
                clickScreen()
            }
    }
    })
}

function clickIdentify() {
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "day",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setMyChartData(data)
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

function clickScreen() {
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "day",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setScreenChartData(data);
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

function initData(){
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "day",
             username:username,
             isAll:isAll
            },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
        },
        success: function (data) {
            setMyChartData(data)
        },
        error: function (data) {
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },

    });
}

if (option && typeof option === "object") {
    myChart.setOption(option, true);
    pieChart.setOption(pieOption, true);
    pPieChart.setOption(pPieOption, true);
    sPieChart.setOption(sPieOption, true);
    nPieChart.setOption(nPieOption, true);
}
var setPieDataNull = function () {

    pieChart.setOption({
        series: [{
            data: [],
        }],
    });
    pPieChart.setOption({
        series: [{
            data: [],
        }],
    });
    sPieChart.setOption({
        series: [{
            data: [],
        }],
    });
    nPieChart.setOption({
        series: [{
            data: [],
        }],
    });
};
var setOptions = function (data) {
    pieChart.setOption({
        title: {
            text: '总量:' + sumTotal(data.total).toLocaleString(),
        },
        color: ['green', 'orange', '#C24641'],
        series: [{
            data: [
                {value: sumTotal(data.n), name: '正常'},
                {value: sumTotal(data.s), name: '性感'},
                {value: sumTotal(data.p), name: '色情'},
            ]
        }],
        legend: {
            data: ['正常', '性感', '色情'],
            formatter: function (name) {
                return formatResult(data, name);
            },
        },
    });
    pPieChart.setOption({
        series: [{
            data: [
                {value: sumTotal(data.p) - sumTotal(data.p_review), name: '确定量'},
                {value: sumTotal(data.p_review), name: '复审量'},
            ]
        }],
    });
    sPieChart.setOption({
        series: [{
            data: [
                {value: sumTotal(data.s) - sumTotal(data.s_review), name: '确定量'},
                {value: sumTotal(data.s_review), name: '复审量'},
            ]
        }],
    });
    nPieChart.setOption({
        series: [{
            data: [
                {value: sumTotal(data.n) - sumTotal(data.n_review), name: '确定量'},
                {value: sumTotal(data.n_review), name: '复审量'},
            ]
        }],
    });
};

var setScreenOptions = function (data) {
    pieChart.setOption({
        title: {
            text: '总量:' + sumTotal(screenTotal(data.isScreen, data.notScreen)).toLocaleString(),
        },
        color: ['#691D93', '#F1D137'],
        series: [{
            name: '拍屏使用统计',
            data: [
                {value: sumTotal(data.isScreen), name: '拍屏'},
                {value: sumTotal(data.notScreen), name: '非拍屏'},
            ]
        }],
        legend: {
            data: ['拍屏', '非拍屏'],
            formatter: function (name) {
                return formatScreenResult(data, name);
            },
        },
    });
};

var setMyChartData = function (data) {   
var v =$("#selectPointOfInterest").val();
var v1 ;
if(v==7||v==1||v==30){
    v1='近'+$("#selectPointOfInterest").val()+'天api调用概况';
}else{
    v1='从'+$("#selectPointOfInterest").val()+'api调用概况';
}
    myChart.setOption({
        
        title: {                            
                text:v1,            
        },
        tooltip: {
            trigger: 'axis',
            textStyle : {
                color: 'white',
                decoration: 'none',
                fontFamily: 'Verdana, sans-serif',
                fontSize: 15,
            },
            formatter: function (params,ticket,callback) {
                var res = "";
                if(data.total.length==0){
                    res += params[0].seriesName+'-<br/>';
                    res += params[1].seriesName+'-<br/>';
                    res += params[2].seriesName+'-<br/>';
                    res += params[3].seriesName+'-'; 
                }else { 
                    for (var i = 0, l = params.length; i < l; i++) {
                        if (i==0)
                        res += data['time'][params[i].dataIndex]+'<br/>';
                        if (i==0||i==3){
                        res += params[i].seriesName+' : '+params[i].value.toLocaleString()+'<br/>';
                        }else if(i==1){
                            res=resSummary(data,params,res,"p",i);
                        }else if(i==2){
                            res=resSummary(data,params,res,"s",i);
                        }
                    }
                }
                callback(ticket, res);
                return res;
            }
        },
        legend: {
            data: ['鉴黄总量','色情','性感','正常']
        },
        color: ['gray','#C24641', 'orange','green'],
        toolbox: {
            show: true,
            feature: {
                dataZoom : {
                    show: true
                },
                dataView: {
                    readOnly: false
                },
                magicType: {
                    type: ['line', 'bar']
                },
                restore: {},
                saveAsImage: {}
            }
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: data.time
        },
        yAxis: {
            axisLabel: {
                formatter: '{value} 次'
            }
        },
        series: [{

            name: '鉴黄总量',
            type: 'line',
            data: data.total,
            barWidth: '20%',
            markLine: {
                data: [{
                    type: 'average',
                    name: '平均值'
                },
                ]
            }
        },{

            name: '色情',
            type: 'line',
            data: data.p,
            barWidth: '20%',
            markLine: {
                data: [{
                    type: 'average',
                    name: '平均值'
                },
                ]
            }
        },{

            name: '性感',
            type: 'line',
            data: data.s,
            barWidth: '20%',
            markLine: {
                data: [{
                    type: 'average',
                    name: '平均值'
                },
                ]
            }
        },{

            name: '正常',
            type: 'line',
            data: data.n,
            barWidth: '20%',
            markLine: {
                data: [{
                    type: 'average',
                    name: '平均值'
                },
                ]
            }
        }],
    },true);
setOptions(data);
};
var setScreenChartData=function(data){
    var v =$("#selectPointOfInterest").val();
    var v1 ;
    if(v==7||v==1||v==30){
        v1='近'+$("#selectPointOfInterest").val()+'天拍屏调用概况';
    }else{
        v1='从'+$("#selectPointOfInterest").val()+'拍屏调用概况';
    }
     myChart.setOption({
        title: {
                text: '近'+$("#selectPointOfInterest").val()+'天拍屏调用概况',            
        },
        tooltip: {
            trigger: 'axis',
            textStyle : {
                color: 'white',
                decoration: 'none',
                fontFamily: 'Verdana, sans-serif',
                fontSize: 15,
            },
            formatter: function (params,ticket,callback) {
                var res = "";
                if(data.isScreen.length==0){
                    res += "拍屏总量"+'-<br/>';
                    res += "拍屏"+'-<br/>';
                    res += "非拍屏"+'-<br/>';
                } else { 
                    for (var i = 0, l = params.length; i < l; i++) {
                        res += data['time'][params[i].dataIndex]+'<br/>';
                        res += params[i].seriesName+' : '+params[i].value.toLocaleString()+'<br/>';
                    }
                }
                callback(ticket, res);
                return res;
            }
        },
        legend: {
            data: ['拍屏总量','拍屏','非拍屏']
        },
        toolbox: {
            show: true,
            feature: {
                dataZoom : {
                    show: true
                },
                dataView: {
                    readOnly: false
                },
                magicType: {
                    type: ['line', 'bar']
                },
                restore: {},
                saveAsImage: {}
            }
        },
        xAxis: {
            type: 'category',
            boundaryGap: false,
            data: data.time
        },
        yAxis: {
            axisLabel: {
                formatter: '{value} 次'
            }
        },
        series: [{
            name: '拍屏总量',
            type: 'line',
            data: screenTotal(data.isScreen,data.notScreen),
            barWidth: '20%',
            markLine: {
                data: [{
                    type: 'average',
                    name: '平均值'
                },
                ]
            }
        },{
            name: '拍屏',
            type: 'line',
            data: data.isScreen,
            barWidth: '20%',
            markLine: {
                data: [{
                    type: 'average',
                    name: '平均值'
                },
                ]
            }
        },{
            name: '非拍屏',
            type: 'line',
            data:data.notScreen,
            barWidth: '20%',
            markLine: {
                data: [{
                    type: 'average',
                    name: '平均值'
                },
                ]
            }
        }],
    },true);
    setScreenOptions(data);
};

$("#selectPointOfInterest").change(function(){
    var value=$("#selectPointOfInterest").val();
    if(curStatus=="identify"){
        if(value==1){
            $('#datepicker').hide()
            curMagnitude="day"
            clickDay();
        }
        else if(value==7){
            $('#datepicker').hide()
            curMagnitude="week"
            clickWeek();
        }
        else if(value==30){
            $('#datepicker').hide()
            curMagnitude="month"
            clickMonth();
        }else{
            now =new Date()
            startTime=addDate(now,-1).format("yyyy-MM-dd"); 
            endTime=now.format("yyyy-MM-dd"); 
            $('#to').datepicker('setStartDate',startTime);
            $("#from").attr("value", startTime)
            $("#to").attr("value", endTime)
            curMagnitude="custom"
            $('#datepicker').show()
        }
    } else{
        if(value==1){
            $('#datepicker').hide()
            curMagnitude="day"
            clickScreenDay();
        }
        else if(value==7){
            $('#datepicker').hide()
            curMagnitude="week"
            clickScreenWeek();
        }
        else if(value==30){
            $('#datepicker').hide()
            curMagnitude="month"
            clickScreenMonth();
        }else{
            now =new Date()
            startTime=addDate(now,-1).format("yyyy-MM-dd"); 
            endTime=now.format("yyyy-MM-dd"); 
            $("#from").attr("value", startTime)
            $("#to").attr("value", endTime)
            curMagnitude="custom"
            $('#datepicker').show()
        }
    }
});

$("#title-identify").click(function () {
    curStatus = "identify"
    $("#p-pie").show()
    $("#s-pie").show()
    $("#n-pie").show()
    if(curMagnitude=="custom"){
        getData()
    }else{
        $.ajax({
            type: 'get',
            url: defaults.chartUrl,
            data: {
                "magnitude": curMagnitude,
                username:username,
                isAll:isAll
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show", {
                    zIndex: 100000,
                });
                setPieDataNull();
            },
            success: function (data) {
                setMyChartData(data);
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },
        });
    }
});

$("#title-screen").click(function () {
    curStatus = "screen"
    $("#p-pie").hide()
    $("#s-pie").hide()
    $("#n-pie").hide()
     if(curMagnitude=="custom"){
        getData()
    }else{
        $.ajax({
            type: 'get',
            url: defaults.chartUrl,
            data: {
                "magnitude": curMagnitude,
                username:username,
                isAll:isAll
            },
            headers: {
                "LoginToken": loginToken,
                "Access-Control-Allow-Headers": "LoginToken",
            },
            beforeSend: function () {
                $.LoadingOverlay("show", {
                    zIndex: 100000,
                });
                setPieDataNull();
            },
            success: function (data) {
                setScreenChartData(data);
            },
            complete: function () {
                $.LoadingOverlay("hide");
            },
            error: function (data) {
            },
        });
    }
});

function clickDay() {
    $("#datepicker").hide();
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "day",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setMyChartData(data)
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

function clickWeek() {
    $("#datepicker").hide();
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "week",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setMyChartData(data)
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

function clickMonth() {
    $("#datepicker").hide();
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "month",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setMyChartData(data);
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

function clickScreenDay() {
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "day",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setScreenChartData(data);
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

function clickScreenWeek() {
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "week",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setScreenChartData(data);
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

function clickScreenMonth() {
    $.ajax({
        type: 'get',
        url: defaults.chartUrl,
        data: {
            "magnitude": "month",
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function () {
            $.LoadingOverlay("show", {
                zIndex: 100000,
            });
            setPieDataNull();
        },
        success: function (data) {
            setScreenChartData(data);
        },
        complete: function () {
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
    });
};

$('#from').datepicker({
    format: "yyyy-mm-dd",
    autoclose : true,  
    endDate : new Date(),
    maxViewMode: 2,
    todayBtn: "linked",
    todayHighlight: true,
}).on('changeDate',function(e){
    startTime = e.date.format("yyyy-MM-dd");
    if(endTime!=undefined){
        getData()
    }
    $('#to').datepicker('setStartDate',startTime);  
});

$('#to').datepicker({
    format: "yyyy-mm-dd",
    autoclose : true,  
    endDate : new Date(),
    maxViewMode: 2,
    todayBtn: "linked",
    todayHighlight: true,
}).on('changeDate',function(e){  
    endTime = e.date.format("yyyy-MM-dd"); 
     if(startTime!=undefined){
        getData()
    }
    $('#from').datepicker('setEndDate',endTime);  
});

function getData(){
    if (curStatus=="identify"){
        $.ajax({
        type: 'get',
        url: defaults.timechartUrl,
        data:{
            "start":startTime,
            "end":endTime,
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        beforeSend: function(){
         $.LoadingOverlay("show",{
            zIndex : 100000,
         });
        setPieDataNull();
        },
         success: function(data) {
           setMyChartData(data);
        },
        complete: function(){
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
        });
    }else{
        $.ajax({
        type: 'get',
        url: defaults.timechartUrl,
        data:{
            "start":startTime,
            "end":endTime,
            username:username,
            isAll:isAll
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
         beforeSend: function(){
         $.LoadingOverlay("show",{
            zIndex : 100000,
         });
         setPieDataNull();
        },
         success: function(data) {
           setScreenChartData(data);
        },
        complete: function(){
            $.LoadingOverlay("hide");
        },
        error: function (data) {
        },
        });
    }   
}   

