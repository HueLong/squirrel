<?php

use Dcat\Admin\Admin;
use Dcat\Admin\Grid;
use Dcat\Admin\Form;
use Dcat\Admin\Grid\Filter;
use Dcat\Admin\Show;

/**
 * Dcat-admin - admin builder based on Laravel.
 * @author jqh <https://github.com/jqhph>
 *
 * Bootstraper for Admin.
 *
 * Here you can remove builtin form field:
 *
 * extend custom field:
 * Dcat\Admin\Form::extend('php', PHPEditor::class);
 * Dcat\Admin\Grid\Column::extend('php', PHPEditor::class);
 * Dcat\Admin\Grid\Filter::extend('php', PHPEditor::class);
 *
 * Or require js and css assets:
 * Admin::css('/packages/prettydocs/css/styles.css');
 * Admin::js('/packages/prettydocs/js/main.js');
 *
 */

//表单默认值
Form::resolving(function (Form $form) {
    $form->disableEditingCheck();
    $form->disableCreatingCheck();
    $form->disableViewCheck();

    $form->tools(function (Form\Tools $tools) {
        $tools->disableDelete();
        $tools->disableView();
    });
});

//列表默认值
Grid::resolving(function (Grid $grid) {
    //排序-默认倒序
    $grid->model()->orderBy($grid->model()->getKeyName(), 'desc');

    //行间距模式
    $grid->tableCollapse(false);

    //默认显示按钮

    //筛选项-默认展开且放置在列表头部
    $grid->filter(function (Grid\Filter $filter) {
        $filter->panel();
        $filter->expand(true);
    });

    $grid->disableBatchDelete();

});
