#ifndef SORTVIEW_H
#define SORTVIEW_H

#include <QWidget>
#include "qlistview.h"
#include "qstringlistmodel.h"
#include "qsortfilterproxymodel.h"
#include "qcombobox.h"
QT_BEGIN_NAMESPACE
namespace Ui { class SortView; }
QT_END_NAMESPACE

class SortView : public QWidget
{
    Q_OBJECT

public:
    SortView(QWidget *parent = nullptr);
    ~SortView();

private:
    Ui::SortView *ui;

    QListView *view;
    QStringListModel *model;
    QSortFilterProxyModel *modelProxy;
    QComboBox *syntaxBox;

private slots:
    void filterChanged(const QString &text);

};
#endif // SORTVIEW_H
