#include "dialog.h"
#include "ui_dialog.h"
#include "qtablewidget.h"
#include "QHBoxLayout"
#include "QItemSelection"
Dialog::Dialog(QWidget *parent)
    : QDialog(parent)
    , ui(new Ui::Dialog)
{
    ui->setupUi(this);

    tb = new QTableWidget(8,4,this);
    auto layout = new QHBoxLayout();
    layout->addWidget(tb);
    this->setLayout(layout);

    auto sm = tb->selectionModel();

    auto topLeft = tb->model()->index(0,0);
    auto bottonRight = tb->model()->index(5,2);

    QItemSelection selection(topLeft,bottonRight);
    sm->select(selection,QItemSelectionModel::Select);

    connect(tb,SIGNAL(selectionChanged(QItemSelection,QItemSelection)),this,SLOT(updateSelection(QItemSelection,QItemSelection)));


}



Dialog::~Dialog()
{
    delete ui;
}

void Dialog::updateSelection(const QItemSelection &selected, const QItemSelection &deselected)
{
    auto model = tb->model();
    QModelIndex index;
    QModelIndexList items = selected.indexes();
    foreach(index,items){
        QString text = QString("(%1,%2)").arg(index.row()).arg(index.column());
        model->setData(index,text);
    }

    items = deselected.indexes();
    foreach(index,items){
        model->setData(index,"");
    }

}

void Dialog::changeCurrent(const QModelIndex &current, const QModelIndex &previous)
{

}































