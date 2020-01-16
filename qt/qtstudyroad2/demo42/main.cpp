

#include <QApplication>
#include "qwidget.h"
#include "qlistwidget.h"
#include "QHBoxLayout"
#include "QLabel"
int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    QWidget window;
    QListWidget* listWidget = new QListWidget();
    listWidget->addItem("chrome");
    listWidget->addItem("ie");
    listWidget->addItem("firefox");
    listWidget->addItem("opera");

    QLabel* label = new QLabel();
    label->setFixedWidth(70);

    QHBoxLayout *layout = new QHBoxLayout;
    layout->addWidget(listWidget);
    layout->addWidget(label);

    window.setLayout(layout);

    window.connect(listWidget,&QListWidget::currentTextChanged,label,&QLabel::setText);
    window.show();

    return a.exec();
}
