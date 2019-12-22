#include <QApplication>
#include "qwidget.h"
#include "qspinbox.h"
#include "qslider.h"
#include "QHBoxLayout"
int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    QWidget window;
    window.setWindowTitle("enter your age");
    QSpinBox *spinbox = new QSpinBox(&window);
    QSlider *slider = new QSlider(Qt::Horizontal,&window);
    spinbox->setRange(0,130);
    slider->setRange(0,130);
    QObject::connect(spinbox,SIGNAL(valueChanged(int)),slider,SLOT(setValue(int)));
    QObject::connect(slider,SIGNAL(valueChanged(int)),spinbox,SLOT(setValue(int)));
    //QObject::connect(spinbox,&QSpinBox::valueChanged,slider,&QSlider::setValue);
    spinbox->setValue(100);
    QHBoxLayout layout;
    layout.addWidget(spinbox);
    layout.addWidget(slider);
    window.setLayout(&layout);


    window.show();
    return a.exec();
}
