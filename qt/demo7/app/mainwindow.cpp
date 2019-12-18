#include "mainwindow.h"
#include "ui_mainwindow.h"
#include "qmessagebox.h"
#include <QToolBar>
MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);
    this->setWindowTitle("MainWindow");
    this->openAction = new QAction(QIcon(":/images/bubble2") , "&Open File",this);
    this->openAction->setShortcut(QKeySequence::Open);
    this->openAction->setStatusTip("Open an existing file.");
    this->connect(this->openAction,&QAction::triggered,this,[=]{this->open();});
    QMenu *menuFile = this->menuBar()->addMenu("&File");
    menuFile->addAction(this->openAction);

    QToolBar *toolbar = this->addToolBar("File");
    toolbar->addAction(this->openAction);

    this->statusBar();

}

MainWindow::~MainWindow()
{
    delete ui;
}

void MainWindow::open()
{
    QMessageBox::information(this,"提示","打开文件");
}
