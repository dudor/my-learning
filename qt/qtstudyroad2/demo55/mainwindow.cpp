#include "mainwindow.h"
#include "qsqldatabase.h"
#include "qmessagebox.h"
#include "QSqlError"
#include "qsqlquery.h"
#include "QVariantList"
#include "qdebug.h"
MainWindow::MainWindow(QWidget *parent)
    : QMainWindow(parent)
{
    if(dbConnect("demo.db")){
        QSqlQuery query;
        query.exec("drop table student");
        if(query.exec("CREATE TABLE STUDENT(ID INTEGER PRIMARY KEY AUTOINCREMENT,NAME VARCHAR,AGE INT)")){
            QMessageBox::information(this,"INFO","CREATE TABLE IS SUCCESS");
        }
        else{
            QMessageBox::critical(this,"ERROR","CREATE TABLE IS FALSE."+ query.lastError().text());
        }

        createData();
    }
}

MainWindow::~MainWindow()
{
}

bool MainWindow::dbConnect(const QString &dbname)
{
    QSqlDatabase db = QSqlDatabase::addDatabase("QSQLITE");
    db.setDatabaseName(dbname);
    if(!db.open()){
        QMessageBox::critical(0,"database error",db.lastError().text());
        return false;
    }
    return true;
}

void MainWindow::createData()
{
    if(dbConnect("demo.db")){
        QSqlQuery query;
        query.prepare("insert into student(name,age) values(?,?)");
        QVariantList names;
        names <<"Tom" << "jack" << "jane" <<"jerry";
        QVariantList ages;
        ages << 20 << 23 << 24 <<25;
        query.addBindValue(names);
        query.addBindValue(ages);

        if(!query.execBatch())
        {
            QMessageBox::critical(this,"error",query.lastError().text());

        }
        query.finish();
        query.exec("select * from student");
        while(query.next())
        {
            QVariant name = query.value(1).toString();
            QVariant age = query.value(2).toString();
            qDebug() << name << age;
        }

    }
    else{
        QMessageBox::critical(this,"ERROR","CONNECT DB IS FALSE.");
    }
}
























