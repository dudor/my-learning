#ifndef FILESYSTEMWIDGET_H
#define FILESYSTEMWIDGET_H

#include <QWidget>
#include "qfilesystemmodel.h"
#include "qtreeview.h"
#include "QHBoxLayout"
#include "QVBoxLayout"
#include "qpushbutton.h"
class FileSystemWidget : public QWidget
{
    Q_OBJECT

public:
    FileSystemWidget(QWidget *parent = nullptr);
    ~FileSystemWidget();

private:
    QFileSystemModel* model;
    QTreeView* treeView;

    void mkdir();
    void rm();

};
#endif // FILESYSTEMWIDGET_H
