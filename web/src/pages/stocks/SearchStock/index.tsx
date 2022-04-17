import React from 'react';
import { Layout } from 'antd';
import styles from './index.less';
import moment from 'moment';
import {
  Form,
  Row,
  Col,
  Input,
  Button,
  Tag,
  Select,
  message,
  Collapse,
  Switch,
  Table,
  Alert,
  Spin,
  Space,
} from 'antd';
import {
  search,
  saveCustomCondition,
  fetchCustomConditionByName,
  deleteCondition,
  updateStockAttr,
  blackList,
} from '@/services/ant-design-pro/search';
import SetColumnValue from './SetColumnValue';

const { Option } = Select;
const { Header, Content } = Layout;
const { Panel } = Collapse;

interface StockSearchProps {
  notifyParent: (value: string[], optionValue: string, isBlock: string) => void;
}

// ====================================================
export async function asyncSaveCustomCondition(param: any): Promise<API.ServiceResponse> {
  try {
    return await saveCustomCondition(param);
  } catch (error) {
    return { success: false, message: '异常，请联系管理员', data: null, total: 0 };
  }
}

export async function asyncBlackList(param: any): Promise<API.ServiceResponse> {
  try {
    return await blackList(param);
  } catch (error) {
    return { success: false, message: '异常，请联系管理员', data: null, total: 0 };
  }
}

export async function asyncUpdateStockAttr(param: any): Promise<API.ServiceResponse> {
  try {
    return await updateStockAttr(param);
  } catch (error) {
    return { success: false, message: '异常，请联系管理员', data: null, total: 0 };
  }
}

export async function asyncDeleteCondition(param: any): Promise<API.ServiceResponse> {
  try {
    return await deleteCondition(param);
  } catch (error) {
    return { success: false, message: '异常，请联系管理员', data: null, total: 0 };
  }
}

export async function asyncSearchStock(param: any): Promise<API.ServiceResponse> {
  try {
    return await search(param);
  } catch (error) {
    return { success: false, message: '异常，请联系管理员', data: null, total: 0 };
  }
}

export async function asyncFetchCustomConditionByName(param: any): Promise<API.ServiceResponse> {
  try {
    return await fetchCustomConditionByName(param);
  } catch (error) {
    return { success: false, message: '异常，请联系管理员', data: null, total: 0 };
  }
}

// ====================================================
const AdvancedSearchForm = (props: any) => {
  const [form] = Form.useForm();

  const deleteConditionA = () => {
    const param = form.getFieldsValue();
    if (param.custom_condtion === undefined || param.custom_condtion === '') return;
    const pp = { name: param.custom_condtion };
    const rsp = asyncDeleteCondition(pp);
    rsp.then((response) => {
      if (response?.success == true) {
        message.info('删除成功');
      } else {
        message.error(response.message, 10);
      }
    });
  };

  const saveCondition = () => {
    const param = form.getFieldsValue();
    if (false !== param.is_option) {
      param.is_option = true;
    }
    if (true !== param.is_too_high) {
      param.is_too_high = false;
    }

    const rsp = asyncSaveCustomCondition(param);
    rsp.then((response) => {
      if (response?.success == true) {
        message.info('保存成功', 10);
      } else {
        message.error(response.message, 10);
      }
    });
  };

  const customConditionChange = (value: string) => {
    const param = { name: value };
    const rsp = fetchCustomConditionByName(param);
    rsp.then((response) => {
      if (response?.success == true) {
        form.setFieldsValue(response.data);
      } else {
        message.error(response.message, 10);
      }
    });
  };

  const onFinish = (param: any) => {
    if (false !== param.is_option) {
      param.is_option = true;
    }
    if (true !== param.is_too_high) {
      param.is_too_high = false;
    }

    props.notifyParent(param);
  };

  const options = props.customSearchList.map((row) => (
    <Option key={row.key} value={row.name}>
      {row.name}
    </Option>
  ));
  return (
    <Form
      form={form}
      name="advanced_search"
      className="ant-advanced-search-form"
      onFinish={onFinish}
    >
      <Row gutter={24}>
        <Col span={2}>
          <Form.Item name="symbol" label="代码">
            <Input />
          </Form.Item>
        </Col>
        <Col span={2}>
          <Form.Item valuePropName="checked" name="is_option" label="期权">
            <Switch checked={false} defaultChecked={true} />
          </Form.Item>
        </Col>
        <Col span={3}>
          <Form.Item name="search_category" label="分类查询">
            <Select>
              <Option value="BK">板块集合</Option>
              <Option value="Star">明星股</Option>
              <Option value="ETF">ETF</Option>
              <Option value="China">中概股</Option>
              <Option value="Hot">热门</Option>
              <Option value="YesterdayHot">昨日强势股</Option>
            </Select>
          </Form.Item>
        </Col>

        <Col span={3}>
          <Form.Item name="buy_tags" label="持有评价">
            <Select>
              <Option value="StrongHold">强烈持有</Option>
              <Option value="Careful">谨慎持有</Option>
              <Option value="WatchCall">做多观察中</Option>
              <Option value="WatchPut">做空观察中</Option>
              <Option value="Put">做空</Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={3}>
          <Form.Item name="company_tags" label="公司性质">
            <Select>
              <Option value="HighUsageRate">高使用率</Option>
              <Option value="StrongCityMoat">护城河</Option>
              <Option value="HighGrowth">高成长</Option>
              <Option value="EmergingIndustry">新兴行业</Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={2}>
          <Form.Item name="profit_tags" label="财务状况">
            <Select>
              <Option value="GoodGain">盈</Option>
              <Option value="Debt">亏</Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={3}>
          <Form.Item name="market_cap_tags" label="市值">
            <Select>
              <Option value="TenToHundred">10-100亿</Option>
              <Option value="LargerThanHundred">大于100亿</Option>
            </Select>
          </Form.Item>
        </Col>
        <Col span={3}>
          <Form.Item name="trade_amount" label="成交额">
            <Select>
              <Option value="LargerThanM">大于1000万</Option>
              <Option value="LargerThanOne">大于1亿</Option>
              <Option value="LargerThanTen">大于10亿</Option>
              <Option value="LargerThanHundred">大于50亿</Option>
            </Select>
          </Form.Item>
        </Col>
        {/*<Col span={4}>
          <Form.Item name="custom_condtion" label="自定义条件">
            <Select onChange={customConditionChange}>{options}</Select>
          </Form.Item>
        </Col>
        <Col span={5}>
          <Form.Item name="custom_name" label="自定义名称">
            <Input />
          </Form.Item>
  </Col>*/}
        <Col span={3} style={{ textAlign: 'right' }}>
          <Button type="primary" htmlType="submit">
            查询
          </Button>
          <Button
            style={{ marginLeft: 8 }}
            onClick={() => {
              form.resetFields();
            }}
          >
            重置
          </Button>
          {/*<Button
            type="primary"
            style={{ marginLeft: 8 }}
            onClick={() => {
              saveCondition();
            }}
          >
            保存
          </Button>
          <Button
            style={{ marginLeft: 8 }}
            onClick={() => {
              deleteConditionA();
            }}
          >
            删除
          </Button>*/}
        </Col>
      </Row>
    </Form>
  );
};
// ====================================================
const columns = [
  {
    title: '代码',
    dataIndex: 'symbol',
    fixed: 'left',
    width: '100px',
    render: (text, record) => {
      let color = '';
      if (record.buy_tags == 'Watch') {
        color = 'gray';
      }
      if (record.buy_tags == 'Careful') {
        color = 'orange';
      }
      if (record.buy_tags == 'StrongHold') {
        color = 'red';
      }
      if (color == '') {
        return <span>{text}</span>;
      } else {
        return (
          <Tag color={color} key={record.id}>
            {text}
          </Tag>
        );
      }
    },
  },
  {
    title: '股票名称',
    dataIndex: 'name',
    width: '180px',
    ellipsis: true,
  },
  {
    title: '市值(亿)',
    dataIndex: 'total_market_cap',
    width: '100px',
    sorter: (a: number, b: number) => a.total_market_cap - b.total_market_cap,
  },
  {
    title: '今日涨',
    dataIndex: 'increase_rate_curr_day',
    width: '100px',
    sorter: (a: number, b: number) => a.increase_rate_curr_day - b.increase_rate_curr_day,
  },
  {
    title: '5日涨',
    dataIndex: 'increase_rate_5day',
    width: '100px',
    sorter: (a: number, b: number) => a.increase_rate_5day - b.increase_rate_5day,
  },
  {
    title: '10日',
    dataIndex: 'increase_rate_10day',
    width: '100px',
    sorter: (a: number, b: number) => a.increase_rate_10day - b.increase_rate_10day,
  },
  {
    title: '20日',
    dataIndex: 'increase_rate_20day',
    width: '100px',
    sorter: (a: number, b: number) => a.increase_rate_20day - b.increase_rate_20day,
  },
  {
    title: '60日',
    dataIndex: 'increase_rate_60day',
    width: '100px',
    sorter: (a: number, b: number) => a.increase_rate_60day - b.increase_rate_60day,
  },
  {
    title: '250日',
    dataIndex: 'increase_rate_250day',
    width: '100px',
    sorter: (a: number, b: number) => a.increase_rate_250day - b.increase_rate_250day,
  },
  {
    title: '当前价',
    dataIndex: 'current_price',
    width: '100px',
    sorter: (a: number, b: number) => a.current_price - b.current_price,
  },
  {
    title: 'PE',
    dataIndex: 'pe',
    width: '100px',
    sorter: (a: number, b: number) => a.pe - b.pe,
  },
  {
    title: '成交(万)',
    dataIndex: 'trade_amount',
    width: '100px',
    sorter: (a: number, b: number) => a.trade_amount - b.trade_amount,
  },
  {
    title: '5日涨额',
    dataIndex: 'trade_rate_5day',
    width: '100px',
    sorter: (a: number, b: number) => a.trade_rate_5day - b.trade_rate_5day,
  },
  {
    title: '10日额',
    dataIndex: 'trade_rate_10day',
    width: '100px',
    sorter: (a: number, b: number) => a.trade_rate_10day - b.trade_rate_10day,
  },
  {
    title: '20日额',
    dataIndex: 'trade_rate_20day',
    width: '100px',
    sorter: (a: number, b: number) => a.trade_rate_20day - b.trade_rate_20day,
  },
  {
    title: '60日额',
    dataIndex: 'trade_rate_60day',
    width: '100px',
    sorter: (a: number, b: number) => a.trade_rate_60day - b.trade_rate_60day,
  },
  {
    title: '成交额时间',
    dataIndex: 'trade_rate_date',
    width: '110px',
    render: (val) => {
      if (val == null) {
        return (<span></span>)
      }
      return (<span>{moment(val).utc().format('YYYY-MM-DD')}</span>)
    },
    
    ellipsis: true,
  },
  {
    title: '行业',
    dataIndex: 'industry',
    width: '140px',
    ellipsis: true,
  },
  {
    title: '更新时间',
    dataIndex: 'update_time',
    width: '110px',
    render: (val) => <span>{moment(val).utc().format('YYYY-MM-DD')}</span>,
    ellipsis: true,
  },
  {
    title: '描述',
    dataIndex: 'description',
    ellipsis: true,
  },
];

// ====================================================
class SearchStock extends React.Component<StockSearchProps> {
  state = {
    selectedRowKeys: [], // Check here to configure the default column
    loading: false,
    data: [],
    customSearchList: [],
    rowId: '',
    optionValue: '', // 对比时的指数名称
  };
  onRef = (ref) => {
    this.child = ref;
  };

  componentDidMount() {
    // const rsp = asyncFetchCustomConditionName()
    // rsp.then((response) => {
    //     if (response?.code == 0) {
    //         if (response.data != null) {
    //             this.setState({ customSearchList: response.data })
    //         }
    //     } else {
    //         message.error(response?.message, 10)
    //     }
    // })
  }

  onSelectChange = (selectedRowKeys: any) => {
    this.setState({ selectedRowKeys });
  };

  setColumnValueCallback = (param: any) => {
    if (this.state.selectedRowKeys.length == 0) {
      message.error('select at lease one row first');
      return;
    }
    param.ids = this.state.selectedRowKeys;
    if (param.buy_tags === undefined) {
      param.buy_tags = '';
    }
    if (param.company_tags === undefined) {
      param.company_tags = '';
    }
    if (param.profit_tags === undefined) {
      param.profit_tags = '';
    }
    if (param.description === undefined) {
      param.description = '';
    }

    const rsp = asyncUpdateStockAttr(param);
    rsp.then((response) => {
      if (response?.success == true) {
        message.info('更新成功', 5);
      } else {
        message.error(response?.message, 10);
      }
    });
  };

  // 模糊查询搜索
  receiveSearchNotifyFromChild = (param: any) => {
    this.setState({
      loading: true,
    });
    const rsp = asyncSearchStock(param);
    rsp
      .then((response) => {
        if (response?.success == true) {
          if (response.data != null) {
            this.setState({ data: response.data });
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

  setRowClassName = (record) => {
    return record.id === this.state.rowId ? styles.clickRowStyl : '';
  };

  redirectToKLine = () => {
    this.props.notifyParent(this.state.selectedRowKeys, this.state.optionValue, 'N');
  };

  redirectToKLineWithBlock = () => {
    this.props.notifyParent(this.state.selectedRowKeys, this.state.optionValue, 'Y');
  };

  handleOptionChange = (value: string) => {
    this.setState({ optionValue: value });
  };

  setBlackList = () => {
    const rsp = asyncBlackList(this.state.selectedRowKeys);
    rsp.then((response) => {
      if (response?.success == true) {
        message.info('更新成功', 5);
      } else {
        message.error(response?.message, 10);
      }
    });
  };

  render() {
    const { selectedRowKeys, loading } = this.state;
    const rowSelection = {
      selectedRowKeys,
      onChange: this.onSelectChange,
    };

    let loadObj;
    if (loading == true) {
      loadObj = (
        <Spin tip="Loading...">
          <Alert message="数据加载中" description="." type="info" />
        </Spin>
      );
    }
    return (
      <Layout>
        <Content>
          <Header className={styles.header}>
            <AdvancedSearchForm
              customSearchList={this.state.customSearchList}
              notifyParent={this.receiveSearchNotifyFromChild}
            />
          </Header>
          <Layout className={styles.bg} style={{ padding: '24px 0' }}>
            <Row>
              <Col>
                <Space>
                  <SetColumnValue notifyParent={this.setColumnValueCallback} />
                </Space>
              </Col>
              <Space style={{ paddingLeft: '5px' }}>
                <Button type="primary" danger onClick={this.setBlackList}>
                  设置黑名单
                </Button>
              </Space>
              <Col>
                <Space style={{ paddingLeft: '5px' }}>
                  <Select className={styles.selectWidth} onChange={this.handleOptionChange}>
                    <Option value="ALL">所有</Option>
                    <Option value="QQQ">纳值</Option>
                    <Option value="DIA">道琼斯指数</Option>
                    <Option value="SPY">标普</Option>
                  </Select>
                  <Button type="primary" onClick={this.redirectToKLine}>
                    对比
                  </Button>
                  <Button type="primary" onClick={this.redirectToKLineWithBlock}>
                    对比模块
                  </Button>
                </Space>
              </Col>
            </Row>
            {loadObj}
            <div>
              <Table
                pagination={{ pageSize: 10 }}
                rowSelection={rowSelection}
                rowClassName={this.setRowClassName}
                columns={columns}
                scroll={{ x: 1300 }}
                onRow={(record) => {
                  return {
                    onClick: (event) => {
                      this.setState({ rowId: record.id });
                      //this.child.childMethod(record?.id, record?.symbol);
                    }, // 点击行
                  };
                }}
                dataSource={this.state.data}
              />
            </div>
          </Layout>
        </Content>
      </Layout>
    );
  }
}

export default SearchStock;
