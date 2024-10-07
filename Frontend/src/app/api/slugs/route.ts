// app/api/slugs/route.ts
import { NextResponse } from 'next/server';
import slugs from '@/data/slugs_data.json';

export async function GET() {
  // 返回 JSON 数据
  return NextResponse.json(slugs);
}
