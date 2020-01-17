#ifndef MYLISTVIEW_H
#define MYLISTVIEW_H

#include <QDialog>
#include "qlistview.h"
#include "qstringlistmodel.h"
class MyListView : public QDialog
{
    Q_OBJECT

public:
    MyListView(QWidget *parent = nullptr);
    ~MyListView();

private:
    void insertData();
    void deleteData();
    void showData();
    QListView* listview;
    QStringListModel* model;

};
#endif // MYLISTVIEW_H
