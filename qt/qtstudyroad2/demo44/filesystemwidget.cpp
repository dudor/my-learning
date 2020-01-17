#include "filesystemwidget.h"
#include "qfilesystemmodel.h"
#include "qtreeview.h"
#include "QHBoxLayout"
#include "QVBoxLayout"
#include "qpushbutton.h"
#include "qinputdialog.h"
#include "qmessagebox.h"
#include "QtDebug"
FileSystemWidget::FileSystemWidget(QWidget *parent)
    : QWidget(parent)
{
    model = new QFileSystemModel;
    model->setRootPath(QDir::currentPath());

    treeView = new QTreeView(this);
    treeView->setModel(model);
    treeView->setRootIndex(model->index(QDir::currentPath()));

    auto layoutBtn = new QHBoxLayout;
    auto btnMkdir = new QPushButton("Mkdir",this);
    auto btnRm = new QPushButton("Rm",this);
    layoutBtn->addWidget(btnMkdir);
    layoutBtn->addWidget(btnRm);

    connect(btnMkdir,&QPushButton::clicked,this,&FileSystemWidget::mkdir);
    connect(btnRm,&QPushButton::clicked,this,&FileSystemWidget::rm);

    auto layout = new QVBoxLayout;
    layout->addWidget(treeView);
    layout->addLayout(layoutBtn);

    this->setLayout(layout);
    setWindowTitle("FileSystemWidget");


}

FileSystemWidget::~FileSystemWidget()
{
}

void FileSystemWidget::mkdir()
{
    auto index = treeView->currentIndex();
    if(!index.isValid()){
        return;
    }
    QString dirName = QInputDialog::getText(this,"create dir","dir name");
    if(!dirName.isEmpty()){
        qDebug()<< index << dirName;
        if(!model->mkdir(index,dirName).isValid()){

            QMessageBox::information(this,"mkdir error","Failed to create the directory");
        }
    }
}

void FileSystemWidget::rm()
{
    auto index = treeView->currentIndex();
    if(!index.isValid()){
        return;
    }
    bool ok;
    if(model->fileInfo(index).isDir()){
        ok = model->rmdir(index);

    }else{
        ok = model->remove(index);
    }
    if(!ok){
        QMessageBox::information(this,"remove",tr("failed to remove %1").arg(model->fileName(index)));
    }



}

