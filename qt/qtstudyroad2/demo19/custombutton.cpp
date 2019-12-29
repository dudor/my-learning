#include "custombutton.h"
#include "qdebug.h"
#include "QMouseEvent"
CustomButton::CustomButton(QWidget *parent) : QPushButton(parent)
{
    this->connect(this,&CustomButton::clicked,this,&CustomButton::onbuttonclicked);
}

void CustomButton::mousePressEvent(QMouseEvent *event)
{
    if(event->button() == Qt::LeftButton){
        qDebug()<< "left";
    }else{
        QPushButton::mousePressEvent(event);
    }
}

void CustomButton::onbuttonclicked()
{
    qDebug() << "CustomButton::onbuttonclicked";
}
