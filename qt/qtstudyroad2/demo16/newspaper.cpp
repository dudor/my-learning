#include "newspaper.h"

Newspaper::Newspaper(QObject *parent) : QObject(parent)
{

}

Newspaper::Newspaper(const QString &name):m_name(name){

}

void Newspaper::send() const
{
    emit newPaper(m_name);
}

