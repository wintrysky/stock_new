import React from 'react';
import { Line } from '@ant-design/charts';
import { message ,Button} from 'antd';
import { getHistoryBySymbol } from '@/services/ant-design-pro/search';

export async function asyncGetHistoryBySymbol(param: any): Promise<API.ServiceResponse> {
  try {
    return await getHistoryBySymbol(param);
  } catch (error) {
    return { success: false, message: '异常，请联系管理员', data: null, total: 0 };
  }
}

interface KLineProps {
  selectedRowKeys: string[];
  optionValue: string;
  isBlock: string;
}

class KLine extends React.Component<KLineProps> {
  state = {
    data :[],
    lineConfig: {data:[]},
  }

  setConfig =(result) => {
     let config = {
      height: 600, // 画布高度
      // padding:画布周围空白，如果字体过长过大，都要填充大一点，auto在苹果设备上可能会出现x轴坐标文字显示不完整
      padding: [40, 0, 40, 60],
      // data 数据源  xField的值取data中某个字段，表示x轴显示的文字。 yField同理 y轴的文字
      data: result.data,
      xField: 'day',
      yField: 'price',
      // seriesField 这个是多条曲线的关键，如果数值有多种，就会出现多条曲线
      seriesField: 'symbol',
      // 设置y轴的样式
      yAxis: {
        line: { style: { stroke: '#0A122E' } }, // 配上这条数据才会显示y轴 stroke等同css color
        // label 配置y轴文字的样式
        label: {
          // formatter 对y轴文字进一步处理
          formatter: (v: any) => `${v}`.replace(/\d{1,3}(?=(\d{3})+$)/g, (s) => `${s},`),
          style: {
            stroke: '#0A122E',
            fontSize: 12,
            fontWeight: 300,
            fontFamily: 'Apercu',
          },
        },
        // grid 配置水平线的样式 下面配置为虚线如果要为实线，不用配置
        grid: {
          line: {
            style: {
              stroke: 'rgb(150,160,171)',
              lineDash: [4, 5],
            },
          },
        },
      },
      xAxis: {
        line: { style: { stroke: '#0A122E' } },
        label: {
          style: {
            stroke: '#0A122E',
            fontSize: 12,
            fontWeight: 300,
            fontFamily: 'Apercu',
          },
        },
      },
      //renderer 画布渲染配置，canvas 或 svg
      renderer: 'svg',
      // 是否为平滑曲线
      smooth: false,
      // 配置显示的2条曲线线条颜色，如果多条，继续添加，注意与右上角的图例颜色要对应
      color: result.colors,
      // 配置显示图例，就是上面那个可以点击的曲线
      legend: {
        //  图例相对于画布的位置
        position: 'top-right',
        // 每个图例的样式
        items: result.items,
      },
    };

    this.setState({lineConfig:config})
  }

  onCompare = () => {
    let param = {};
    param.compare_list = this.props.selectedRowKeys;
    param.option_name = this.props.optionValue;
    param.is_block = this.props.isBlock;
    const rsp = asyncGetHistoryBySymbol(param);
    rsp
      .then((response) => {
        if (response?.success == true) {
          if (response.data != null) {
            this.setState({data: response.data.data});
            this.setConfig(response.data)
          }else{
            this.setState({data: []});
            this.setConfig([])
          }
        } else {
          message.error(response?.message, 10);
        }
      })
      .catch(() => {
        message.error('查询错误', 10);
      })
      .finally(() => {
        this.setState({
          selectedRowKeys: [],
          loading: false,
        });
      });
  };

  componentDidMount = () => {
    console.log('', this.props.selectedRowKeys, this.props.optionValue);
  };

  render() {
    return (
      <div>
        <Button type="primary" onClick={this.onCompare}>
          开始对比
        </Button>
        <Line height={700} {...this.state.lineConfig} />
      </div>
    );
  }
}

export default KLine;
