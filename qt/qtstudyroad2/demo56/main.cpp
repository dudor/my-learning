#include <QCoreApplication>
#include "qsqldatabase.h"
#include "qsqlquery.h"
#include "QSqlQuery"
#include "qmessagebox.h"
#include "qsqlerror.h"
#include "qsqltablemodel.h"
#include "qsqlrecord.h"
#include "qdebug.h"
#include "qtableview.h"
#include "qapplication.h"
#include "qheaderview.h"
#include "qsqlrelationaltablemodel.h"
#include "qsqlrelationaldelegate.h"
bool dbConnect(const QString &dbname)
{
    QSqlDatabase db = QSqlDatabase::addDatabase("QSQLITE");
    db.setDatabaseName(dbname);
    if(!db.open()){
        QMessageBox::critical(0,"database error",db.lastError().text());
        return false;
    }
    return true;
}

void getData(){
    if(dbConnect("demo.db"))
    {
        QSqlTableModel model;
        model.setTable("student");
        model.setFilter("age > 23");
        if(model.select()){
            for ( int i = 0;i<model.rowCount();i++)
            {
                QSqlRecord row = model.record(i);
                QString name = row.value("name").toString();
                int age = row.value("age").toInt();
                qDebug() << name << age;
            }
        }
    }
}
void addData(){
    QSqlTableModel model;
    model.setTable("student");
    int row = 0;
    model.insertRows(row,1);
    model.setData(model.index(row,1),"chen");
    model.setData(model.index(row,2),26);
    model.submitAll();
}
void showTable(){
    if(dbConnect("demo.db")){
        QSqlRelationalTableModel *model= new QSqlRelationalTableModel;
        model->setTable("student");
        model->setSort(1,Qt::AscendingOrder);
        model->setHeaderData(1,Qt::Horizontal,"Name");
        model->setHeaderData(2,Qt::Horizontal,"AGE");
        model->setHeaderData(3,Qt::Horizontal,"CITY");
        model->setRelation(3,QSqlRelation("CITY","id","name"));
        model->select();

        QTableView *view = new QTableView;
        view->setModel(model);
        view->setSelectionMode(QAbstractItemView::SingleSelection);
        view->setSelectionBehavior(QAbstractItemView::SelectRows);
        view->resizeRowsToContents();
        //view->setEditTriggers(QAbstractItemView::NoEditTriggers);
        view->setItemDelegate(new QSqlRelationalDelegate(view));
        QHeaderView *header = view->horizontalHeader();
        header->setStretchLastSection(true);
        view->show();
    }
}

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    //getData();
//    addData();
    getData();
    showTable();
    return a.exec();
}
