@extends('layout')

@section('header')
    <div class="page-header">
        <h1><i class="glyphicon glyphicon-edit"></i> ItemLogs / Edit #{{$item_log->id}}</h1>
    </div>
@endsection

@section('content')
    @include('error')

    <div class="row">
        <div class="col-md-12">

            <form action="{{ route('item_logs.update', $item_log->id) }}" method="POST">
                <input type="hidden" name="_method" value="PUT">
                <input type="hidden" name="_token" value="{{ csrf_token() }}">

                <div class="form-group @if($errors->has('action')) has-error @endif">
                       <label for="action-field">Action</label>
                    <input type="text" id="action-field" name="action" class="form-control" value="{{ $item_log->action }}"/>
                       @if($errors->has("action"))
                        <span class="help-block">{{ $errors->first("action") }}</span>
                       @endif
                    </div>
                    <div class="form-group @if($errors->has('item_id')) has-error @endif">
                       <label for="item_id-field">Item_id</label>
                    <input type="text" id="item_id-field" name="item_id" class="form-control" value="{{ $item_log->item_id }}"/>
                       @if($errors->has("item_id"))
                        <span class="help-block">{{ $errors->first("item_id") }}</span>
                       @endif
                    </div>
                <div class="well well-sm">
                    <button type="submit" class="btn btn-primary">Save</button>
                    <a class="btn btn-link pull-right" href="{{ route('item_logs.index') }}"><i class="glyphicon glyphicon-backward"></i>  Back</a>
                </div>
            </form>

        </div>
    </div>
@endsection