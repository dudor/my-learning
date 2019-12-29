#include <QCoreApplication>
#include "qfile.h"
#include "qdatastream.h"
#include "qdebug.h"

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);

    QFile file("file.dat");
    if(!file.exists()){
        file.open(QIODevice::ReadWrite|QIODevice::NewOnly);
    }else{
        file.open(QIODevice::ReadWrite);
    }
    QDataStream out(&file);
    out << QString("the answer is ");
    out << qint32(42);
    file.close();

    file.open(QIODevice::ReadOnly);
    QString str;
    qint32 ab;
    QDataStream in(&file);
    in >> str >> ab;
    file.close();
    qDebug() << str << ab;


    return a.exec();
}
