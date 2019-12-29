#include "widget.h"
#include "ui_widget.h"
#include "qpaintengine.h"
#include "qpainter.h"
#include "QPaintEvent"
#include "qpen.h"
Widget::Widget(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::Widget)
{
    ui->setupUi(this);
    this->resize(800,800);
}

Widget::~Widget()
{
    delete ui;
}

void Widget::paintEvent(QPaintEvent *event)
{
    QPainter painter(this);
    painter.drawLine(100,100,200,200);
    painter.setPen(Qt::red);
    painter.drawRect(110,110,220,220);
    painter.setPen(QPen(Qt::green,5));
    painter.setBrush(Qt::blue);
    painter.drawEllipse(330,330,440,440);


    painter.setPen(QPen(Qt::black, 5, Qt::DashDotLine, Qt::RoundCap));
    painter.setBrush(Qt::yellow);
    painter.drawEllipse(50, 150, 200, 150);

    painter.setRenderHint(QPainter::Antialiasing, true);
    painter.setPen(QPen(Qt::black, 5, Qt::DashDotLine, Qt::RoundCap));
    painter.setBrush(Qt::yellow);
    painter.drawEllipse(300, 150, 200, 150);
}


