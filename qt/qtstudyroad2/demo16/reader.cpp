#include "reader.h"
#include "qdebug.h"
Reader::Reader(QObject *parent) : QObject(parent)
{

}



void Reader::receiveNewspaper(const QString &name) const
{
    qDebug() << "receiveNewspaper " << name;
}
