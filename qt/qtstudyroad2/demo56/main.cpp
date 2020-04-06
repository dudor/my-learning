#include <QCoreApplication>
#include "qsqldatabase.h"
#include "qsqlquery.h"
#include "QSqlQuery"
#include "qmessagebox.h"
#include "qsqlerror.h"
#include "qsqltablemodel.h"
int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    return a.exec();
}

void getData(){
if(dbConnect("demo.db"))
{

}

}

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
