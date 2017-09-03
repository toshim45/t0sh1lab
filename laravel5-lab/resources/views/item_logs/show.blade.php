@extends('layout')
@section('header')
<div class="page-header">
        <h1>ItemLogs / Show #{{$item_log->id}}</h1>
        <form action="{{ route('item_logs.destroy', $item_log->id) }}" method="POST" style="display: inline;" onsubmit="if(confirm('Delete? Are you sure?')) { return true } else {return false };">
            <input type="hidden" name="_method" value="DELETE">
            <input type="hidden" name="_token" value="{{ csrf_token() }}">
            <div class="btn-group pull-right" role="group" aria-label="...">
                <a class="btn btn-warning btn-group" role="group" href="{{ route('item_logs.edit', $item_log->id) }}"><i class="glyphicon glyphicon-edit"></i> Edit</a>
                <button type="submit" class="btn btn-danger">Delete <i class="glyphicon glyphicon-trash"></i></button>
            </div>
        </form>
    </div>
@endsection

@section('content')
    <div class="row">
        <div class="col-md-12">

            <form action="#">
                <div class="form-group">
                    <label for="nome">ID</label>
                    <p class="form-control-static"></p>
                </div>
                <div class="form-group">
                     <label for="action">ACTION</label>
                     <p class="form-control-static">{{$item_log->action}}</p>
                </div>
                    <div class="form-group">
                     <label for="item_id">ITEM_ID</label>
                     <p class="form-control-static">{{$item_log->item_id}}</p>
                </div>
            </form>

            <a class="btn btn-link" href="{{ route('item_logs.index') }}"><i class="glyphicon glyphicon-backward"></i>  Back</a>

        </div>
    </div>

@endsection