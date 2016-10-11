/**
 * Created by huang on 2016/7/28.
 */
 $(document).ready(function() {
var defaults = {
    picUrl: "/api/photo_rate",
    roomUrl: "/api/room_rate",
    missUrl: "/api/miss_rate",
    searchUrl:"/stat_search_users",
};
var dom = document.getElementById("picreview");
var myChart = echarts.init(dom);
var app = {};
var username=""
var curRate="photo"
if (getUrlParam("username") == null) {
 username ="yk"
} else {
 username = getUrlParam("username");
}
function getUrlParam(name) {
         var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
         var r = window.location.search.substr(1).match(reg);  //匹配目标参数
         if (r != null) return unescape(r[2]);
         return null; //返回参数值
     }

$.ajax({
    type: 'get',
    url: defaults.picUrl,
    data:{
            username:username
        },
    headers: {
        "LoginToken": loginToken,
        "Access-Control-Allow-Headers": "LoginToken",
    },

    success: function(data) {
        data.rate = _.map(data.rate, function (v) {
                toPercent(v)
                return v
            });
        data.rate_inke = _.map(data.rate_inke, function (v) {
            toPercent(v)
            return v
        });
        myChart.setOption({
             title: {
            text: '实时图片复审率',
            },
            tooltip: {
                trigger: 'axis'
            },
            legend: {
                data: ['复审率','inke复审率']
            },
            toolbox: {
                show: true,
                feature: {
                    dataZoom: {

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
                    formatter: '{value} '
                }
            },
            dataZoom: [

                {
                    type: 'inside',
                    xAxisIndex: [0],

                },
                {
                    type: 'inside',
                    yAxisIndex: [0],

                }
            ],
            series: [{
                name: '复审率',
                type: 'line',
                data: data.rate,
                barWidth: '20%',

                markLine: {
                    data: [{
                        type: 'average',
                        name: '平均值'
                    },

                    ]
                }
            },{
                 name: 'inke复审率',
                type: 'line',
                data: data.rate_inke,
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
    },
    error: function (data) {
    },

});

function getRoomRate(){
    var username = $("#autocomplete-search").val();
    if (getUrlParam("username") == null) {
        username ="yk"
    } else {
        username = getUrlParam("username");
    }
    $.ajax({
        type: 'get',
        url: defaults.roomUrl,
        data:{
            username: username,
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
            var legend=[]
            var series=[]
            data.forEach(function(item,idx){
                legend.push(item.version)
                var result={
                    name: item.version,
                    type: 'line',
                    data: item.result.rate,
                    barWidth: '20%',
                    markLine: {
                        data: [{
                            type: 'average',
                            name: '平均值'
                        },
                        ]
                    }
                }
                series.push(result)
            })
             myChart.setOption({
                title: {
                text: '实时房间复审率',
                },
                tooltip: {
                    trigger: 'axis'
                },
                legend: {
                    data: legend
                },
                toolbox: {
                    show: true,
                    feature: {
                        dataZoom: {

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
                    data: data[0].result.time
                },
                yAxis: {
                    axisLabel: {
                        formatter: '{value} '
                    }
                },
                dataZoom: [

                    {
                        type: 'inside',
                        xAxisIndex: [0],

                    },
                    {
                        type: 'inside',
                        yAxisIndex: [0],

                    }
                ],
                series:series,
            },true);
        },
        error: function (data) {
        },
    });
}
var toPercent = function (v) {
    return (v * 100).toFixed(4) + "%"
};
function getPhotoRate(){
    var username = $("#autocomplete-search").val();
    if (getUrlParam("username") == null) {
        username ="yk"
    } else {
        username = getUrlParam("username");
    }
    $.ajax({
        type: 'get',
        url: defaults.picUrl,
        data:{
            username:username
        },
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
              data.rate = _.map(data.rate, function (v) {
                toPercent(v)
                return v
            });
             data.rate_inke = _.map(data.rate_inke, function (v) {
                toPercent(v)
                return v
            });
             myChart.setOption({
             title: {
                text: '实时图片复审率',
                },
                tooltip: {
                    trigger: 'axis'
                },
                legend: {
                    data: ['复审率','inke复审率']
                },
                toolbox: {
                    show: true,
                    feature: {
                        dataZoom: {

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
                        formatter: '{value} '
                    }
                },
                dataZoom: [

                    {
                        type: 'inside',
                        xAxisIndex: [0],

                    },
                    {
                        type: 'inside',
                        yAxisIndex: [0],

                    }
                ],
                series: [{

                    name: '复审率',
                    type: 'line',
                    data: data.rate,
                    barWidth: '20%',

                    markLine: {
                        data: [{
                            type: 'average',
                            name: '平均值'
                        },

                        ]
                    }
                },{
                 name: 'inke复审率',
                type: 'line',
                data: data.rate_inke,
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
        },
        error: function (data) {
        },
    });

}

function getMissRate(){
    var username = $("#autocomplete-search").val();
    if (getUrlParam("username") == null) {
        username ="yk"
    } else {
        username = getUrlParam("username");
    }
    $.ajax({
        type: 'get',
        url: defaults.missUrl,
        headers: {
            "LoginToken": loginToken,
            "Access-Control-Allow-Headers": "LoginToken",
        },
        success: function(data) {
            var legend=[]
            var series=[]
            var time=[]
            if (data.length>0)
               time=data[0].result['time']
            data.forEach(function(item,idx){
                if(time.length<item.result['time'].length)
                    time=item.result['time']
                var name=""
                if(item.type=="min"){
                    name="拉黑第一帧_"+item.version
                }else{
                    name="拉黑最后一帧_"+item.version
                }
                legend.push(name)
                var value=0
                var miss=0
                var total=0
                 item.result.miss.forEach(function(item,idx){
                    miss=miss+item
                 });
                 item.result.total.forEach(function(item,idx){
                    total=total+item
                 });
                value=(miss/total).toFixed(4)
                var result={
                    name: name,
                    type: 'line',
                    data: item.result.rate,
                    barWidth: '20%',
                    markLine: {
                        data: [  
                             [
                             {value: value, coord:[0,(Math.max.apply(null,item.result.rate)+Math.min.apply(null,item.result.rate))/2.05]}, 
                             {coord:[ item.result.time.length-1,(Math.max.apply(null,item.result.rate)+Math.min.apply(null,item.result.rate))/2.05],
                                name: '平均值',
                            },     
                            ]
                        ]
                    },
                }
                series.push(result)
            })
             myChart.setOption({
                title: {
                text: '漏检率',
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
                        if(data.length==0){
                            res += "拉黑第一帧_v6_2"+'-<br/>';
                            res += "拉黑第一帧_v6_3"+'-<br/>';
                            res += "拉黑最后一帧_v6_2"+'-<br/>';
                            res += "拉黑最后一帧_v6_3"+'-<br/>';
                        } else { 
                            for (var i = 0, l = params.length; i < l; i++) {
                                if(i==0){
                                    res += data[0].result['time'][params[i].dataIndex]+'<br/>';
                                }
                                res += params[i].seriesName+' : '+params[i].value+'<br/>';
                                res += "漏检数"+' : '+data[i].result["miss"][params[i].dataIndex]+'<br/>';
                                res += "总数"+' : '+data[i].result["total"][params[i].dataIndex]+'<br/>';
                            }
                        }
                        callback(ticket, res);
                        return res;
                    }
                },
                legend: {
                    data: legend
                },
                toolbox: {
                    show: true,
                    feature: {
                        dataZoom: {

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

                xAxis: [
                    {
                        type: 'category',
                        boundaryGap: false,
                        data: data[0].result.time
                    }
                ],
                yAxis: [
                    {
                        type:"value",
                        axisLabel: {
                            formatter: '{value} '
                        }
                    }
                ],
                series:series,
            },true);
        },
        error: function (data) {
        },
    });
}
$("#title-room").click(function(data){
    curRate="room"
    getRoomRate()
});
$("#title-pic").click(function(){
    curRate="photo"
    getPhotoRate()
});
$("#title-miss").click(function(){
    curRate="miss"
    getMissRate()
});
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
            window.location.href="/searchuser?url=picreview&text="+$("#autocomplete-search").val()
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
            if(curRate=="photo")
            {
               getPhotoRate()
            }else if(curRate=="room"){
               getRoomRate()
            }else{
               getMissRate()
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
});
