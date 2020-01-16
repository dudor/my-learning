#ifndef MYLISTVIEW_H
#define MYLISTVIEW_H

#include <QDialog>

class MyListView : public QDialog
{
    Q_OBJECT

public:
    MyListView(QWidget *parent = nullptr);
    ~MyListView();
};
#endif // MYLISTVIEW_H
