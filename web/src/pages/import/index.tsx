import React from 'react';
import { Card, Select } from 'antd';
import styles from './index.less';
import UploadDrag from './UploadDrag';

const { Option } = Select;

class ImportExcel extends React.Component {
  state = {
    selectedType: 'all',
    dateString: '',
  };

  componentDidMount() {}

  handleChange = (value: string) => {
    this.setState({ selectedType: encodeURI(value) });
  };

  onDateChange = (date, dateString) => {
    this.setState({ dateString: encodeURI(dateString) });
  };

  render() {
    return (
      <div className={styles.container}>
        <Select defaultValue="all" style={{ width: 300 }} onChange={this.handleChange} showSearch>
          <Option value="all">所有美股</Option>
          <Option value="bk">板块集合</Option>
          <Option value="option">可校验期权</Option>
          <Option value="china">中概股</Option>
          <Option value="ETF">ETF</Option>
          <Option value="hot">热门股票</Option>
          <Option value="star">明星股</Option>
          <Option value="yestodayhot">昨日强势股</Option>
        </Select>
        {/*<DatePicker onChange={this.onDateChange} />*/}
        <Card bordered={true}>
          <div className={styles.container}>
            <UploadDrag importType={this.state.selectedType} dateString={this.state.dateString} />
          </div>
        </Card>
      </div>
    );
  }
}

export default ImportExcel;
