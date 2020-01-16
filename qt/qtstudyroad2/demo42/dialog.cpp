#include "dialog.h"
#include "qlabel.h"
#include "qlistwidget.h"
#include "QHBoxLayout"
#include "qtreewidget.h"
#include "qtablewidget.h"
#include "qtablewidget.h"

Dialog::Dialog(QWidget *parent)
    : QDialog(parent)
{
    TableWidgetDemo();
}

Dialog::~Dialog()
{
}

void Dialog::ListWidgetDemo()
{
    auto label = new QLabel(this);
    label->setFixedWidth(70);
    label->setText("demo42");
    auto listwidget = new QListWidget(this);
    listwidget->addItem("IE");
    listwidget->addItem("FIREFOX");
    listwidget->addItem("CHROME");
    listwidget->addItem("OPERA");
    listwidget->addItem("THEWORLD");
    listwidget->addItem("MAXTHON");

    auto layout = new QHBoxLayout(this);
    layout->addWidget(label);
    layout->addWidget(listwidget);

    this->setLayout(layout);

    connect(listwidget, SIGNAL(currentTextChanged(QString)),
            label, SLOT(setText(QString)));
}

void Dialog::TreeWidgetDemo()
{
    auto treewidget = new QTreeWidget(this);
    treewidget->setColumnCount(2);
    treewidget->setHeaderLabels(QStringList()<<"NAME"<<"LEVEL");
    auto root = new QTreeWidgetItem(treewidget,QStringList("root")<<"1");
    root->addChild(new QTreeWidgetItem(QStringList("leaf1")<<"2"));
    root->addChild(new QTreeWidgetItem(QStringList("leaf2")<<"2"));

    treewidget->addTopLevelItem(root);

    treewidget->show();

}

void Dialog::TableWidgetDemo()
{
    this->resize(500,500);
    auto tablewidget = new QTableWidget(this);
    tablewidget->setColumnCount(4);
    tablewidget->setRowCount(5);
    QStringList headers;
    headers << "ID"<<"NAME"<<"AGE"<<"SEX";
    tablewidget->setHorizontalHeaderLabels(headers);
    tablewidget->setItem(0,0,new QTableWidgetItem("0001"));
    tablewidget->setMaximumWidth(500);
    tablewidget->setMaximumHeight(500);
    tablewidget->show();
}

