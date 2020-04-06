#include "sortview.h"
#include "ui_sortview.h"
#include "qlineedit.h"
#include "qlabel.h"
#include "QHBoxLayout"
#include "QVBoxLayout"
#include "qcombobox.h"
#include "qregexp.h"
SortView::SortView(QWidget *parent)
    : QWidget(parent)
    , ui(new Ui::SortView)
{
    ui->setupUi(this);

    model = new QStringListModel(QColor::colorNames(),this);

    modelProxy = new QSortFilterProxyModel(this);
    modelProxy->setSourceModel(model);
    modelProxy->setFilterKeyColumn(0);

    view = new QListView(this);
    view->setModel(modelProxy);

    QLineEdit *filterInput = new QLineEdit;
    QLabel *filterLabel = new QLabel("Filter:");
    QHBoxLayout *filterLayout = new QHBoxLayout;
    filterLayout->addWidget(filterLabel);
    filterLayout->addWidget(filterInput);

    syntaxBox = new QComboBox;
    syntaxBox->setSizePolicy(QSizePolicy::Expanding,QSizePolicy::Preferred);
    syntaxBox->addItem("Regular expression",QRegExp::RegExp);
    syntaxBox->addItem("Wildcard",QRegExp::Wildcard);
    syntaxBox->addItem("Fixed string",QRegExp::FixedString);
    QLabel *syntaxLabel = new QLabel("Syntax");
    QHBoxLayout *syntaxLayout = new QHBoxLayout;
    syntaxLayout->addWidget(syntaxLabel);
    syntaxLayout->addWidget(syntaxBox);

    QVBoxLayout *mainLayout = new QVBoxLayout(this);
    mainLayout->addWidget(view);
    mainLayout->addLayout(filterLayout);
    mainLayout->addLayout(syntaxLayout);

    connect(filterInput,SIGNAL(textChanged(QString)),this,SLOT(filterChanged(QString)));





}

SortView::~SortView()
{
    delete ui;
}

void SortView::filterChanged(const QString &text)
{
    QRegExp::PatternSyntax syntax = QRegExp::PatternSyntax(syntaxBox->itemData(syntaxBox->currentIndex()).toInt());
    QRegExp regexp(text,Qt::CaseInsensitive, syntax);
    modelProxy->setFilterRegExp(regexp);


}

