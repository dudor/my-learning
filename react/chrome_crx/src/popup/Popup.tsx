import React, { useState } from 'react'
import { Card, Avatar, Radio, Image, Row, Badge, Col, Space } from 'antd';
import { EditOutlined, EllipsisOutlined, SettingOutlined, UserOutlined } from '@ant-design/icons';
const { Meta } = Card;


const Popup: React.FC = () => {
    const [prov, setProv] = useState('')
    const options = [
        { label: 'Apple', value: 'Apple' },
        { label: 'Pear', value: 'Pear' },
        { label: 'Orange', value: 'Orange' },
    ];
    return (
        <Card
            style={{ width: 300 }}
            actions={[
                <SettingOutlined key="setting" />,
                <EditOutlined key="edit" />,
                <EllipsisOutlined key="ellipsis" />,
            ]}
        >
            {/* <Meta
                avatar={<Avatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />}
                title="Card title"
                description="This is the description"
            /> */}
                <Row justify="center">
                    <Col>
                        <span className="avatar-item">
                            <Badge count={1}>
                                <Avatar size={64} icon={<UserOutlined />} />
                            </Badge>
                        </span>
                    </Col>
                </Row>
                <Row justify="center" style={{margin:5}}>
                    <Col>
                        <Radio.Group
                            options={options}
                            // onChange={onChange3}
                            // value={value3}
                            size="small"
                            optionType="button"
                            buttonStyle="solid"
                        />
                    </Col>
                </Row>
        </Card>
    )
}

export default Popup;
