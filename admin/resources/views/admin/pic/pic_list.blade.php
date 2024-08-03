<style>
    .pic-list {
        width: 80%;
        margin-left: 10%;
        text-align: center;
        border-style: dashed;
        border-width: 1px;
    }
</style>

<table class="pic-list">
    <tr style="height: 50px;">
        <th>ID</th>
        <th>名称</th>
        <th>图片</th>
    </tr>
    @foreach($list as $pic)
        <tr style="height: 50px;border-style: dashed;border-width: 1px;">
            <td>{{$pic['id']}}</td>
            <td>{{$pic['name']}}</td>
            <td><img src="{{get_img_link($pic['pic_url'][0]??"",'s')}}"  alt="封面图"/></td>
        </tr>
    @endforeach
</table>
