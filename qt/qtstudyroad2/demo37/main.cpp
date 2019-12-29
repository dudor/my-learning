#include <QCoreApplication>
#include "qfile.h"
#include "qtextstream.h"
#include "qdebug.h"

int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);
    QFile file("file.txt");
    if(file.open(QFile::WriteOnly|QFile::Truncate|QFile::NewOnly)){
        QTextStream in(&file);
        in << "the answer is " << 42;
    }

    file.close();

    QTextStream out(&file);
    file.open(QFile::ReadOnly);
    qDebug() << out.readAll();


    return a.exec();
}
