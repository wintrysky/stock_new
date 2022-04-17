import React from 'react';
import { Layout, Tabs } from 'antd';
import KLine from './KLine';
import SearchStock from './SearchStock';

const { TabPane } = Tabs;

class Stocks extends React.Component {
  state = {
    activeKey: 'search',
    selectedRowKeys: [],
    optionValue: '',
    isBlock: 'N',
  };

  componentDidMount = () => {};

  onChange = (activeKey) => {
    this.setState({ activeKey });
  };

  receiveSearchNotifyFromChild = (bastList: string[],optionValue: string,isBlock: string) => {
    console.log("----",bastList,optionValue)
    this.setState({
      activeKey: 'compare',
      selectedRowKeys: bastList,
      optionValue: optionValue,
      isBlock:isBlock,
    });
  };

  render() {
    const { activeKey } = this.state;
    return (
      <Layout>
        <Tabs type="card" onChange={this.onChange} activeKey={activeKey} tabBarGutter={4}>
          <TabPane tab={'查询数据'} key={'search'} closable={false}>
            <SearchStock notifyParent={this.receiveSearchNotifyFromChild} />
          </TabPane>
          <TabPane tab={'对比图'} key={'compare'} closable={true}>
            <KLine selectedRowKeys={this.state.selectedRowKeys} optionValue={this.state.optionValue} isBlock={this.state.isBlock} />
          </TabPane>
        </Tabs>
      </Layout>
    );
  }
}

export default Stocks;
