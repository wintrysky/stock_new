import React from 'react';
import { Upload, message } from 'antd';
import styles from './index.less';
import { InboxOutlined } from '@ant-design/icons';
const { Dragger } = Upload;

const props = {
  name: 'file',
  multiple: false,
  withCredentials: false,
  accept: '.csv',
  action: 'service/import/csv'
};

interface CustomProps {
  importType: string,
  dateString: string,
}

// 上传附件控件
class Uploaddrag extends React.Component<CustomProps> {
  state = {
    uploadType: false, // false:上传文件,true:上传文件夹
    uploadFullPath: '', // 保存上传文件的全路径，包括文件夹名称
    folders: [],
    alerts: [], // 告警列表文字信息
    defaultFileList: [], // 编辑模式下，默认已上传的附件,
    showFolderLoading: false, // 显示文件夹加载中
  }

  // 改变上传方式：文件，文件夹
  onUploadTypeChange = (e: any) => {
    this.setState({ uploadType: e.target.value });
  }

  // 上传控件上传文件时,将文件夹全路径传递到headers中
  draggerOnChangeFunc = (info: any) => {
    const { status } = info.file;
    if (status === 'done') {
      const { response } = info.file;
      if (response.success === true) {
        message.success(`${info.file.name} 文件上传成功.`);
      } else {
        message.error(`失败：${response.message}`, 12);
      }
    } else if (status === 'error') {
      message.error(`${info.file.name} 失败:${info.file.response?.message}`, 12);
    }
  }

  render() {
    return (
      <div className={styles.container}>
        <div id="components-upload-demo-drag">
          <Dragger defaultFileList={this.state.defaultFileList}
            headers={{ importtype: this.props.importType,dateString:this.props.dateString }} {...props} onChange={this.draggerOnChangeFunc}
            ref={(ref) => this.myDragger = ref}> 
            <p className="ant-upload-drag-icon">
              <InboxOutlined />
            </p>
            <p className="ant-upload-text">
              点击或拖拽目标到此区域上传，支持文件夹或多个文件
            </p>
          </Dragger>
        </div>
      </div>);
  }
}

export default Uploaddrag;
