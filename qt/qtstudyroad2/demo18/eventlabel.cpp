#include "eventlabel.h"
#include "qlabel.h"
#include "QMouseEvent"
void eventlabel::mouseMoveEvent(QMouseEvent *ev)
{

    this->setText(QString("<center><h1>Move: (%1, %2)</h1></center>")
                    .arg(QString::number(ev->x()), QString::number(ev->y())));
}

void eventlabel::mousePressEvent(QMouseEvent *ev)
{
    this->setText(QString("<center><h1>Press: (%1, %2)</h1></center>")
                  .arg(QString::number(ev->x()), QString::number(ev->y())));
}
void eventlabel::mouseReleaseEvent(QMouseEvent *ev)
{
    QString msg;
    msg.sprintf("<center><h1>Release: (%d, %d)</h1></center>",
                ev->x(), ev->y());
    this->setText(msg);
}
