#include <QApplication>
#include "qlabel.h"
#include "custombutton.h"
int main(int argc, char *argv[])
{
    QApplication a(argc, argv);


    CustomButton btn;
    btn.setText("点我");
    btn.show();
    return a.exec();
}
