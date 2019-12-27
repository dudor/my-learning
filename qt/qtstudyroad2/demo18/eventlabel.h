#ifndef EVENTLABEL_H
#define EVENTLABEL_H

#include <QObject>
#include "qlabel.h"
class eventlabel : public QLabel
{
    Q_OBJECT
protected:
    void mouseMoveEvent(QMouseEvent *ev) override;
    void mousePressEvent(QMouseEvent *ev) override;
    void mouseReleaseEvent(QMouseEvent *ev) override;
public:


signals:

};

#endif // EVENTLABEL_H
