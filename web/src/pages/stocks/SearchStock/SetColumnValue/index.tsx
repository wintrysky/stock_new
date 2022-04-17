import styles from './index.less';
import React, { useState } from 'react';
import { Button, Modal, Form, Input, Select } from 'antd';
const { Option } = Select;

const layout = {
  wrapperCol: {
    span: 24,
  },
};

interface Values {
  title: string;
  description: string;
  modifier: string;
}


interface CollectionCreateFormProps {
  visible: boolean;
  onCreate: (values: Values) => void;
  onCancel: () => void;
}

interface CollectionsPageProps {
  notifyParent: (value: any) => void;
}

const CollectionCreateForm: React.FC<CollectionCreateFormProps> = ({
  visible,
  onCreate,
  onCancel,
}) => {
  const [form] = Form.useForm();

  return (
    <Modal
      visible={visible} title="设置描述"
      okText="提交"
      cancelText="取消"
      onCancel={onCancel}
      onOk={() => {
        form
          .validateFields()
          .then((values: any) => {
            form.resetFields();
            onCreate(values);
          })
          .catch((info) => {
            console.log('Validate Failed:', info);
          });
      }}
    >
      <Form
        form={form}
        {...layout}
        layout="horizontal"
        name="form_in_modal"
        initialValues={{ modifier: 'public' }}
      ><Form.Item name="buy_tags" label="持有评价">
          <Select allowClear={true}>
            <Option value="--">--</Option>
            <Option value="StrongHold">强烈持有</Option>
            <Option value="Careful">谨慎持有</Option>
            <Option value="WatchCall">做多观察中</Option>
            <Option value="WatchPut">做空观察中</Option>
            <Option value="Put">做空</Option>
            <Option value="BlackList">黑名单</Option>
          </Select>
        </Form.Item>
        <Form.Item name="company_tags" label="公司性质">
          <Select allowClear={true}>
            <Option value="--">--</Option>
            <Option value="HighUsageRate">高使用率</Option>
            <Option value="StrongCityMoat">护城河</Option>
            <Option value="HighGrowth">高成长</Option>
            <Option value="EmergingIndustry">新兴行业</Option>
          </Select>
        </Form.Item>
        <Form.Item name="profit_tags" label="财务状况">
          <Select>
            <Option value="--">--</Option>
            <Option value="GoodGain">盈利良好</Option>
            <Option value="Debt">亏损</Option>
          </Select>
        </Form.Item>
        <Form.Item name="description" label="描述">
          <Input.TextArea rows={12} />
        </Form.Item>

      </Form>
    </Modal>
  );
};

const CollectionsPage = (props: CollectionsPageProps) => {
  const [visible, setVisible] = useState(false);

  // 提交
  const onCreate = (values: any) => {
    props.notifyParent(values);
    setVisible(false);
  };

  return (
    <div>
      <Button id="db_owner_btn"
        onClick={() => {
          setVisible(true);
        }}
      >
        设置属性
      </Button>
      <CollectionCreateForm
        visible={visible}
        onCreate={onCreate}
        onCancel={() => {
          setVisible(false);
        }}
      />
    </div>
  );
};

export default (props: CollectionsPageProps) => (
  <div className={styles.container}>
    <div id="components-form-demo-form-in-modal">
      <CollectionsPage {...props} />
    </div>
  </div>
);
