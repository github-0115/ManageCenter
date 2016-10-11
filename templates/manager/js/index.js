
var defaults = {
    chartUrl: "/console_detail",
};
var dom = document.getElementById("daily");
var myChart = echarts.init(dom);
var pPieChart = echarts.init(document.getElementById('p-pie'));
var sPieChart = echarts.init(document.getElementById('s-pie'));
var nPieChart = echarts.init(document.getElementById('n-pie'));
var app = {};
var cnResult= {
    p:"色情",
    s:"性感",
    n:"正常",
}

option = {
    tooltip: {
        trigger: 'axis'
    },
    legend: {
        data: ['总次数']
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
        data: []
    },
    yAxis: {
        axisLabel: {
            formatter: '{value} 次'
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
        name: '处理数量',
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
pPieOption = {
    title: [{
        text: '色情',
        x: 'center',
        y: 'center',
        itemGap: 20,
        textStyle : {
            color : 'rgba(30,144,255,0.8)',
            fontFamily : '微软雅黑',
            fontSize : 15,
            fontWeight : 'bolder'
        }
    },
    ],

    tooltip : {
        show: true,
        formatter: formatPieTip
    },
    color:['#f23434','#c22a80'],
    series : [
        {
            name:'色情复审占比',
            type:'pie',
            clockWise:false,
            radius : [25, 45],
            itemStyle:{
                normal: {
                    label:{
                        formatter: formatLabel
                    },
                    labelLine: {length:0}
                },
            },
        }
    ]
};

sPieOption = {
    title:[{
        text: '性感',
        x: 'center',
        y: 'center',
        itemGap: 20,
        textStyle : {
            color : 'rgba(30,144,255,0.8)',
            fontFamily : '微软雅黑',
            fontSize : 15,
            fontWeight : 'bolder'
        }
    },
    ],
    tooltip : {
        show: true,
        formatter: formatPieTip
    },
    color:['orange','#c22a80'],
    series : [
        {
            name:'性感复审占比',
            type:'pie',
            clockWise:false,
            radius : [25, 45],
            itemStyle:{
                normal: {
                    label:{
                        formatter: formatLabel
                    },
                    labelLine: {length:0}
                },
            },
        }
    ]
};

nPieOption = {
    title: [{
        text: '正常',
        x: 'center',
        y: 'center',
        itemGap: 20,
        textStyle : {
            color : 'rgba(30,144,255,0.8)',
            fontFamily : '微软雅黑',
            fontSize : 15,
            fontWeight : 'bolder'
        }

    },
    ],
    tooltip : {
        show: true,
        formatter: formatPieTip
    },
    color:['green','#c22a80'],
    series : [
        {
            name:'正常复审占比',
            type:'pie',
            clockWise:false,
            radius : [25, 45],
            itemStyle:{
                normal: {
                    label:{
                        formatter: formatLabel
                    },
                    labelLine: {length:0}
                },
            },
        }
    ]
};

function formatLabel(params){
    var res = "";
    res=params.name+" : "+params.value.toLocaleString();
    return res;
};
function formatPieTip (params) {
    var res = "";
    res=params.seriesName+"<br/>"+params.name+" : "+params.value.toLocaleString()+" ("+params.percent+"%)"
    return res;
};
var resSummary = function (data,params,res,type,idx) {
    res+=cnResult[type]+':'+data[type][params[idx].dataIndex].toLocaleString()+'&nbsp;&nbsp;&nbsp;&nbsp;'+'(需复审：'+data[type+'_review'][params[idx].dataIndex].toLocaleString()+'）<br/>';
    return res
};
var setOptions=function(data){
     //$("#total").text(data.today_detail.today_total).toLocaleString();
    pPieChart.setOption({
        series : [{
            data:[
                {value:data.today_detail.p-data.today_detail.p_review, name:'确定量'},
                {value:data.today_detail.p_review, name:'复审量'},
            ],
            itemStyle:{
                normal: {
                    label:{
                        formatter: function(params){
                            var res = "";
                            /*if(params.name=="总量")
                             res=params.name+" : "+sumTotal(data.total).toLocaleString();
                             else*/
                            res=params.name+" : "+params.value.toLocaleString();
                            return res;
                        }
                    },

                },
            },
        }],
    });
    sPieChart.setOption({
        series : [{
            data:[
                {value:data.today_detail.s-data.today_detail.s_review, name:'确定量'},
                {value:data.today_detail.s_review, name:'复审量'},
            ],
            itemStyle:{
                normal: {
                    label:{
                        formatter: function(params){
                            var res = "";
                            /*if(params.name=="总量")
                             res=params.name+" : "+sumTotal(data.total).toLocaleString()
                             else*/
                            res=params.name+" : "+params.value.toLocaleString();
                            return res;
                        }
                    },

                },
            },
        }],
    });
    nPieChart.setOption({
        series : [{
            data:[
                {value:data.today_detail.n-data.today_detail.n_review, name:'确定量'},
                {value:data.today_detail.n_review, name:'复审量'},
            ],
            itemStyle:{
                normal: {
                    label:{
                        formatter: function(params){
                            var res = "";
                            /* if(params.name=="总量")
                             res=params.name+" : "+sumTotal(data.total).toLocaleString()
                             else*/
                            res=params.name+" : "+params.value.toLocaleString();
                            return res;
                        }
                    },

                },
            },
        }],
    });
};
$.ajax({
    type: 'get',
    url: defaults.chartUrl,
    headers: {
        "LoginToken": loginToken,
        "Access-Control-Allow-Headers": "LoginToken",
    },
    beforeSend: function(){
        $.LoadingOverlay("show",{
            zIndex : 100000,
        });

    },
    success: function(data) {
        Increasepercent=(data.today_detail.today_total-data.today_detail.yesterday_total)/data.today_detail.yesterday_total
        reviewtotal=data.today_detail.n_review +data.today_detail.s_review+data.today_detail.p_review
        reviewrate=(data.today_detail.n_review +data.today_detail.s_review+data.today_detail.p_review)/data.today_detail.today_total
        if(data.today_detail.yesterday_total==0){
            Increasepercent=0
        }

        if(data.today_detail.today_total==0){
            reviewrate = 0
        }


        $(".Increasepercent").text((Increasepercent*100).toFixed(2)+"%")
        $("#total").text(data.today_detail.today_total.toLocaleString())
        $("#reviewtotal").text(reviewtotal.toLocaleString())
        $("#newuser").text(data.submit_user)
        $("#staruser").text(data.startUser)
        $(".reviewrate").text("（总复审率"+((reviewrate*100).toFixed(2))+"%）")
         myChart.setOption({
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
                    if(data.today_rep.total.length==0){
                        res += "今天"+'-<br/>';

                    }
                    if(data.yesterday_rep.total.length==0){
                        res += "昨天"+'-<br/>';
                    }
                    if(data.today_rep.total.length!=0||data.yesterday_rep.total.length!=0)
                        for (var i = 0, l = params.length; i < l; i++) {
                            if(params[i].value!=null){
								if (i%2==0){
									res += data.yesterday_rep['time'][params[i].dataIndex]+'<br/>';
								}                               
                                res += params[i].seriesName+' : '+params[i].value.toLocaleString()+'<br/>';
                            }
                    }
                    callback(ticket, res);
                    return res;
                }
            },
            legend: {
                data: ['今天','昨天']
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
                data: data.yesterday_rep.time
            },
            yAxis: {
                axisLabel: {
                    formatter: '{value} 次'
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
                name: '今天',
                type: 'line',
                data: data.today_rep.total,
                barWidth: '20%',
                markLine: {
                    data: [{
                        type: 'average',
                        name: '平均值'
                    },
                    ]
                }
            },{
                name: '昨天',
                type: 'line',
                data: data.yesterday_rep.total,
                barWidth: '20%',
                markLine: {
                    data: [{
                        type: 'average',
                        name: '平均值'
                    },
                    ]
                }
            },
            ],
        },true);
        setOptions(data);
    },
    complete: function(){
        $.LoadingOverlay("hide");
    },
    error: function (data) {
    },
});


if (option && typeof option === "object") {
    myChart.setOption(option, true);
    pPieChart.setOption(pPieOption,true);
    sPieChart.setOption(sPieOption,true);
    nPieChart.setOption(nPieOption,true);
}
