#include "mylistview.h"
#include "QHBoxLayout"
#include "qstringlist.h"
#include "qstringlistmodel.h"
#include "qlistview.h"
#include "qpushbutton.h"
#include "QVBoxLayout"
#include "qinputdialog.h"
#include "qmessagebox.h"
#include "qdebug.h"
MyListView::MyListView(QWidget *parent)
    : QDialog(parent)
{
    QStringList data;
    data.append("LETTER A");
    data << "LETTER B" << "LETTER C";

    model = new QStringListModel(this);
    model->setStringList(data);

    listview = new QListView(this);
    listview->setModel(model);

    auto btnLayout = new QHBoxLayout;
    auto btnInsert = new QPushButton("insert",this);
    connect(btnInsert, &QPushButton::clicked, this, &MyListView::insertData);
    auto btnDelete = new QPushButton("delete",this);
    connect(btnDelete, &QPushButton::clicked, this, &MyListView::deleteData);
    auto btnShow = new QPushButton("show",this);
    connect(btnShow, &QPushButton::clicked, this, &MyListView::showData);
    btnLayout->addWidget(btnInsert);
    btnLayout->addWidget(btnDelete);
    btnLayout->addWidget(btnShow);

    auto mainLayout = new QVBoxLayout;
    mainLayout->addWidget(listview);
    mainLayout->addLayout(btnLayout);
    this->setLayout(mainLayout);


}

MyListView::~MyListView()
{
}

void MyListView::insertData()
{

    bool isOk;
    QString text = QInputDialog::getText(this,"insert","please text new data",QLineEdit::Normal,"you are insert new data",&isOk);
    if(isOk){
        /*
        auto row = listview->currentIndex().row();
        model->insertRows(row+1,1);
        QModelIndex index = model->index(row+1);
        model->setData(index,text);
        listview->setCurrentIndex(index);
        listview->edit(index);
        */

        auto index = listview->currentIndex();
        auto row = index.row();
        qDebug()<< index << row;
        if(row < 0){
            row = 0;

        }
        model->insertRows(row,1);
        model->setData(index,text);
        listview->edit(index);
    }
}

void MyListView::deleteData()
{
    if(model->rowCount()>1){
        model->removeRows(listview->currentIndex().row(),1);
    }
}

void MyListView::showData()
{
    QStringList data = model->stringList();
    QString str;
    foreach (QString item , data) {
        str += item + "\n";
    }
    QMessageBox::information(this,"show data",str);
}

