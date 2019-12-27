#include <QApplication>
#include "eventlabel.h"
int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    eventlabel label;
    label.setText("mouse event demo");
    label.resize(200,200);
    label.show();

    return a.exec();
}
