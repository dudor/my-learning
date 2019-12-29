#include "mainwindow.h"

#include <QApplication>
#include "qfile.h"
#include "qdebug.h"
#include "qfileinfo.h"
#include "qdir.h"
#include "qdatastream.h"
int main(int argc, char *argv[])
{
    QApplication a(argc, argv);

    QFile file("in.txt");
    if(!file.open(QIODevice::ReadOnly | QIODevice::Text)){
        qDebug()<< "Open file failed.";
    }else{
        while (!file.atEnd()) {
            qDebug() << file.readLine().toStdString().c_str();
        }
    }
    file.close();

    QFileInfo fileinfo(file);
    qDebug()<< fileinfo.isDir();
    qDebug()<< fileinfo.isExecutable();
    qDebug()<< fileinfo.baseName();
    qDebug()<< fileinfo.completeBaseName();
    qDebug()<< fileinfo.suffix();
    qDebug()<< fileinfo.completeSuffix();
    qDebug()<< fileinfo.filePath();
    qDebug()<< fileinfo.absoluteFilePath();
    qDebug()<< fileinfo.canonicalFilePath();
    qDebug()<< QDir::currentPath();




    return a.exec();
}
